package token

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claimsEnvelope struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type jwtAccessTokenService struct {
	key           []byte
	signingMethod jwt.SigningMethod
}

func NewJwtAccessTokenService(key []byte) AccessTokenService {
	return &jwtAccessTokenService{
		key: key,
	}
}

func (s *jwtAccessTokenService) Create(ctx context.Context, claims *AccessTokenClaims) (string, error) {
	c := claimsEnvelope{
		claims.UserID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
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

func (s *jwtAccessTokenService) Validate(ctx context.Context, tokenString string) (*AccessTokenClaims, error) {
	var envelope claimsEnvelope
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

	return &AccessTokenClaims{UserID: envelope.UserID}, nil
}
