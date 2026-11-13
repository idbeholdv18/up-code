package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/idbeholdv18/up-code/app/hash"
	"github.com/idbeholdv18/up-code/app/token"
	"github.com/idbeholdv18/up-code/app/user"
)

type httpHandler struct {
	userService        user.Service
	accessTokenService token.AccessTokenService
	hashService        hash.HashService
}

func NewHttpHandler(userService user.Service, accessTokenService token.AccessTokenService, hashService hash.HashService) *httpHandler {
	return &httpHandler{
		userService:        userService,
		accessTokenService: accessTokenService,
		hashService:        hashService,
	}
}

func (h *httpHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var body RegisterRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("incorrect req body")
		return
	}
	defer r.Body.Close()

	if body.Email == "" || body.FullName == "" || len(body.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("incorrect req body")
		return
	}

	hashedPassword, err := h.hashService.Hash(r.Context(), body.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("internal error %v", err))
		return
	}

	u, err := h.userService.Create(r.Context(), body.Email, hashedPassword, body.FullName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("internal error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RegisterResponsePayload{
		Email:    u.Email,
		FullName: u.FullName,
	})
}

func (h *httpHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var body LoginRequestBody
	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("incorrect req body")
		return
	}
	defer r.Body.Close()

	if body.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("incorrect req body")
		return
	}

	candidate, err := h.userService.FindByEmail(r.Context(), body.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("incorrect email or password")
		return
	}

	isPasswordCorrect, err := h.hashService.Compare(r.Context(), body.Password, candidate.PasswordHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("internal error %v", err))
		return
	}

	if !isPasswordCorrect {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("incorrect email or password")
		return
	}

	claims := token.AccessTokenClaims{
		UserID: candidate.ID,
	}

	ss, err := h.accessTokenService.Create(r.Context(), claims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("create token error")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&LoginResponsePayload{
		AccessToken: ss,
	})
}
