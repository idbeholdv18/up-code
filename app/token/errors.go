package token

import (
	"errors"
	"net/http"

	"github.com/idbeholdv18/up-code/shared/errregistry"
)

var (
	ErrInvalidAccessToken      = errors.New("invalid access token")
	ErrInvalidConifmationToken = errors.New("invalid confirmation token")
)

func Errors() *errregistry.Registry {
	registry := errregistry.New()

	registry.RegisterSentinel(ErrInvalidAccessToken, errregistry.APIError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       "INVALID_ACCESS_TOKEN",
		Message:    "token is invalid or expired",
	})

	registry.RegisterSentinel(ErrInvalidConifmationToken, errregistry.APIError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "INVALID_CONFIRMATION_TOKEN",
		Message:    "token is invalid or expired",
	})

	return registry
}
