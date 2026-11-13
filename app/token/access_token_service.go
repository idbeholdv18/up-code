package token

import "context"

type AccessTokenService interface {
	Create(ctx context.Context, claims AccessTokenClaims) (string, error)
	Validate(ctx context.Context, tokenString string) (*AccessTokenClaims, error)
}
