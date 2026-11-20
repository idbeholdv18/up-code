package httpkit

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

const RequestIDHeader = "X-Request-ID"

type reqIDCtxKey struct{}

func WithRequestID(ctx context.Context, reqID string) context.Context {
	return context.WithValue(ctx, reqIDCtxKey{}, reqID)
}

func RequestIDMiddleware() Middleware {
	return func(next Handler) Handler {
		return func(w http.ResponseWriter, r *http.Request) error {
			reqID := r.Header.Get(RequestIDHeader)
			if reqID == "" {
				reqID = GenerateRequestID()
			}
			w.Header().Set(RequestIDHeader, reqID)

			ctx := WithRequestID(r.Context(), reqID)
			return next(w, r.WithContext(ctx))
		}
	}
}

func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(reqIDCtxKey{}).(string); ok {
		return reqID
	}
	return ""
}

func GenerateRequestID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
