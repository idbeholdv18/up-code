package auth

import (
	"encoding/json"
	"net/http"

	"github.com/idbeholdv18/up-code/app/token"
	"github.com/idbeholdv18/up-code/app/user"
)

type httpHandler struct {
	userService        user.Service
	accessTokenService token.AccessTokenService
}

func NewHttpHandler(userService user.Service, accessTokenService token.AccessTokenService) *httpHandler {
	return &httpHandler{
		userService:        userService,
		accessTokenService: accessTokenService,
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

	u, err := h.userService.Create(r.Context(), body.Email, body.Password, body.FullName)
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

	if candidate.PasswordHash != body.Password {
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
