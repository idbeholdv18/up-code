package auth

import (
	"encoding/json"
	"net/http"

	"github.com/idbeholdv18/up-code/app/user"
)

type httpHandler struct {
	userService user.Service
}

func NewHttpHandler(userService user.Service) *httpHandler {
	return &httpHandler{
		userService: userService,
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
