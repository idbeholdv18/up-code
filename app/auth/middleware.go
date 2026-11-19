package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/idbeholdv18/up-code/app/token"
	"github.com/idbeholdv18/up-code/shared/httpkit"
)

type ctxKey string

const userIDKey ctxKey = "user_id"

func GetUserIDFromCtx(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDKey).(int64)
	return id, ok
}

func JwtMiddleware(accessTokenService token.AccessTokenService) httpkit.Middleware {
	return func(next httpkit.Handler) httpkit.Handler {
		return func(w http.ResponseWriter, r *http.Request) error {
			w.Header().Add("Content-Type", "application/json")

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				return ErrInvalidAccessToken
			}

			tokenString, ok := strings.CutPrefix(authHeader, "Bearer ")
			if !ok {
				return ErrInvalidAccessToken
			}

			claims, err := accessTokenService.Validate(r.Context(), tokenString)
			if err != nil {
				return ErrInvalidAccessToken
			}

			ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)

			return next(w, r.WithContext(ctx))
		}
	}
}
