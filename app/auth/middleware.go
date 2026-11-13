package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/idbeholdv18/up-code/app/token"
)

type ctxKey string

const userIDKey ctxKey = "user_id"

func GetUserIDFromCtx(ctx context.Context) (int64, bool) {
	id, ok := ctx.Value(userIDKey).(int64)
	return id, ok
}

func JwtMiddleware(accessTokenService token.AccessTokenService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("missing access token: must be bearer")
				return
			}

			tokenString, ok := strings.CutPrefix(authHeader, "Bearer ")
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode("invalid access token")
				return
			}

			claims, err := accessTokenService.Validate(r.Context(), tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(fmt.Sprintf("invalid access token %v", err))
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
