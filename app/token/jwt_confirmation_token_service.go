package token

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type confirmationTokenClaimsEnvelope struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type jwtConfirmationTokenService struct {
	key           []byte
	signingMethod jwt.SigningMethod
	ttl           time.Duration
}

func NewJwtConfirmationTokenService(key []byte, ttl time.Duration) ConfirmationTokenService {
	return &jwtConfirmationTokenService{
		key:           key,
		signingMethod: jwt.SigningMethodHS256,
		ttl:           ttl,
	}
}

func (s *jwtConfirmationTokenService) Create(ctx context.Context, claims ConfirmationTokenClaims) (string, error) {
	c := confirmationTokenClaimsEnvelope{
		claims.UserID,
		claims.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	ss, err := t.SignedString(s.key)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s *jwtConfirmationTokenService) Parse(ctx context.Context, tokenString string) (*ConfirmationTokenClaims, error) {
	var envelope confirmationTokenClaimsEnvelope
	token, err := jwt.ParseWithClaims(tokenString, &envelope, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != s.signingMethod.Alg() {
			return nil, errors.New("incorrect signing method")
		}
		return s.key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token validation failed")
	}

	return &ConfirmationTokenClaims{UserID: envelope.UserID, Email: envelope.Email}, nil
}
