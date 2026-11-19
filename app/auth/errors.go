package auth

import (
	"errors"
	"net/http"

	"github.com/idbeholdv18/up-code/shared/errregistry"
)

var (
	ErrIvalidRegisterRequestBody = errors.New("incorrect register requeset body")
	ErrInvalidLoginRequestBody   = errors.New("incorrect login requeset body")
	ErrUserNotActivated          = errors.New("user not activated")
	ErrIncorrectPasswordOrEmail  = errors.New("incorrect password or email")
	ErrInvalidActivationToken    = errors.New("invalid activation token")
	ErrInvalidAccessToken        = errors.New("invalid access token")
)

func Errors() *errregistry.Registry {
	reg := errregistry.New()

	reg.RegisterSentinel(ErrIvalidRegisterRequestBody, errregistry.APIError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "INVALID_PAYLOAD",
		Message:    "invalid payload",
	})

	reg.RegisterSentinel(ErrInvalidLoginRequestBody, errregistry.APIError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "INVALID_PAYLOAD",
		Message:    "invalid payload",
	})

	reg.RegisterSentinel(ErrUserNotActivated, errregistry.APIError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "NOT_ACTIVATED",
		Message:    "user not activated",
	})

	reg.RegisterSentinel(ErrIncorrectPasswordOrEmail, errregistry.APIError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       "INCORRECT_CREDENTIALS",
		Message:    "user not activated",
	})

	reg.RegisterSentinel(ErrInvalidActivationToken, errregistry.APIError{
		HTTPStatus: http.StatusBadRequest,
		Code:       "INVALID_ACTIVATION_TOKEN",
		Message:    "activation token is invalid",
	})

	reg.RegisterSentinel(ErrInvalidAccessToken, errregistry.APIError{
		HTTPStatus: http.StatusUnauthorized,
		Code:       "INVALID_ACCESS_TOKEN",
		Message:    "access token is invalid",
	})

	return reg
}
