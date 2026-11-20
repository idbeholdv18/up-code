package token

import "context"

type AccessTokenClaims struct {
	UserID int64 `json:"user_id"`
}

type ConfirmationTokenClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
}

type AccessTokenService interface {
	Create(ctx context.Context, claims AccessTokenClaims) (string, error)
	Validate(ctx context.Context, tokenString string) (*AccessTokenClaims, error)
}

type ConfirmationTokenService interface {
	Create(ctx context.Context, claims ConfirmationTokenClaims) (string, error)
	Parse(ctx context.Context, tokenString string) (*ConfirmationTokenClaims, error)
}
