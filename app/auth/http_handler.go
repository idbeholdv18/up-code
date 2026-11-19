package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/idbeholdv18/up-code/app/events"
	"github.com/idbeholdv18/up-code/app/hash"
	"github.com/idbeholdv18/up-code/app/token"
	"github.com/idbeholdv18/up-code/app/user"
)

type httpHandler struct {
	userService              user.Service
	accessTokenService       token.AccessTokenService
	confirmationTokenService token.ConfirmationTokenService
	hashService              hash.HashService
	eventsPublisher          events.Publisher
	baseURL                  string
}

func NewHttpHandler(
	userService user.Service,
	accessTokenService token.AccessTokenService,
	hashService hash.HashService,
	confirmationTokenService token.ConfirmationTokenService,
	eventsPublisher events.Publisher,
	baseURL string,
) *httpHandler {
	return &httpHandler{
		userService:              userService,
		accessTokenService:       accessTokenService,
		confirmationTokenService: confirmationTokenService,
		hashService:              hashService,
		eventsPublisher:          eventsPublisher,
		baseURL:                  baseURL,
	}
}

func (h *httpHandler) HandleRegister(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	var body RegisterRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return ErrIvalidRegisterRequestBody
	}
	defer r.Body.Close()

	if body.Email == "" || body.FullName == "" || len(body.Password) < 8 {
		return ErrIvalidRegisterRequestBody
	}

	hashedPassword, err := h.hashService.Hash(r.Context(), body.Password)
	if err != nil {
		// TODO switch to api error
		log.Printf("failed to hash password: %v", err)
		return err
	}

	u, err := h.userService.Create(r.Context(), body.Email, hashedPassword, body.FullName)
	if err != nil {
		// TODO switch to api error
		log.Printf("failed to create user: %v", err)
		return err
	}

	confirmationToken, err := h.confirmationTokenService.Create(r.Context(), token.ConfirmationTokenClaims{
		UserID: u.ID,
		Email:  u.Email,
	})

	if err != nil {
		// TODO switch to api error
		log.Printf("failed to create confirmation token: %v", err)
		return err
	}

	confirmationLink := fmt.Sprintf("%s/auth/confirm?token=%s", h.baseURL, confirmationToken)

	envelope := events.NewEnvelope(
		events.TopicUserRegistered,
		events.UserRegisteredPayload{
			UserID:           u.ID,
			Email:            u.Email,
			FullName:         u.FullName,
			ConfirmationLink: confirmationLink,
		},
	)

	err = h.eventsPublisher.Publish(r.Context(), events.TopicUserRegistered, envelope)
	if err != nil {
		// TODO fix Dual Write
		// TODO switch to api error
		log.Printf("failed to pubish registration event: %v", err)
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RegisterResponsePayload{
		Email:    u.Email,
		FullName: u.FullName,
	})

	return nil
}

func (h *httpHandler) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	var body LoginRequestBody
	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return ErrInvalidLoginRequestBody
	}
	defer r.Body.Close()

	if body.Email == "" {
		return ErrInvalidLoginRequestBody
	}

	candidate, err := h.userService.FindByEmail(r.Context(), body.Email)
	if err != nil {
		// TODO switch to api error
		log.Printf("failed to find user by email: %v", err)
		return err
	}

	if candidate.ActivatedAt == nil {
		return ErrUserNotActivated
	}

	isPasswordCorrect, err := h.hashService.Compare(r.Context(), body.Password, candidate.PasswordHash)
	if err != nil {
		// TODO switch to api error
		log.Printf("failed to compare password with hash: %v", err)
		return err
	}

	if !isPasswordCorrect {
		return ErrIncorrectPasswordOrEmail
	}

	claims := token.AccessTokenClaims{
		UserID: candidate.ID,
	}

	ss, err := h.accessTokenService.Create(r.Context(), claims)
	if err != nil {
		// TODO switch to api error
		log.Printf("failed to create access token: %v", err)
		return err
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&LoginResponsePayload{
		AccessToken: ss,
	})

	return nil
}

func (h *httpHandler) HandleActivate(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")

	rawToken := r.URL.Query().Get("token")
	if rawToken == "" {
		return ErrInvalidActivationToken
	}

	claims, err := h.confirmationTokenService.Parse(r.Context(), rawToken)
	if err != nil {
		// TODO switch to api error
		log.Printf("failed to parse confirmation token: %v", err)
		return err
	}

	u, err := h.userService.Activate(r.Context(), claims.UserID)
	if err != nil {
		// TODO switch to api error
		log.Printf("failed to activate user: %v", err)
		return err
	}

	// TODO bad check. Move to repo
	if u.ActivatedAt == nil {
		return errors.New("user was not activated in unknown reasons")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ActivateResponsePayload{
		Email:       u.Email,
		FullName:    u.FullName,
		ActivatedAt: *u.ActivatedAt,
	})

	return nil
}
