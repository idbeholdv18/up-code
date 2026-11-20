package user

import (
	"errors"
	"net/http"

	"github.com/idbeholdv18/up-code/shared/errregistry"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidData       = errors.New("invalid data")
)

func Errors() *errregistry.Registry {
	registry := errregistry.New()
	registry.RegisterSentinel(ErrUserAlreadyExists, errregistry.APIError{
		HTTPStatus: http.StatusConflict,
		Code:       "USER_ALREADY_EXISTS",
		Message:    "user with such email already exists",
	})

	registry.RegisterSentinel(ErrInvalidData, errregistry.APIError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "INVALID_DATA",
		Message:    "invalid data",
	})

	return registry
}
