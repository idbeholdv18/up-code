package token

import "context"

type ConfirmationTokenService interface {
	Create(ctx context.Context, claims ConfirmationTokenClaims) (string, error)
	Parse(ctx context.Context, tokenString string) (*ConfirmationTokenClaims, error)
}
