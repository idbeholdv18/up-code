package httpkit

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/idbeholdv18/up-code/shared/errregistry"
)

type Handler func(w http.ResponseWriter, r *http.Request) error
type Middleware func(next Handler) Handler

type RouteBuilder struct {
	errorRegistry *errregistry.Registry
	mws           []Middleware
}

func NewRouteBuilder(reg *errregistry.Registry) *RouteBuilder {
	return &RouteBuilder{
		errorRegistry: reg,
		mws:           []Middleware{},
	}
}

func (b *RouteBuilder) clone() *RouteBuilder {
	clonedMws := make([]Middleware, len(b.mws))
	copy(clonedMws, b.mws)

	return &RouteBuilder{
		errorRegistry: b.errorRegistry,
		mws:           clonedMws,
	}
}

func (b *RouteBuilder) Use(mw Middleware) *RouteBuilder {
	b.mws = append(b.mws, mw)
	return b
}

func (b *RouteBuilder) With(mw Middleware) *RouteBuilder {
	clone := b.clone()
	return clone.Use(mw)
}

func (b *RouteBuilder) Group(fn func(sub *RouteBuilder)) {
	fn(b.clone())
}

func (b *RouteBuilder) Handle(h Handler) http.HandlerFunc {
	errorHandlingWrapper := func(w http.ResponseWriter, r *http.Request) error {
		if err := h(w, r); err != nil {
			apiErr := b.errorRegistry.Resolve(err)
			apiErr.RequestID = GetRequestID(r.Context())

			if apiErr.HTTPStatus >= http.StatusInternalServerError {
				slog.ErrorContext(r.Context(), "internal server error",
					slog.String("code", apiErr.Code),
					slog.String("err_msg", err.Error()),
				)
			} else {
				slog.WarnContext(r.Context(), "request rejected",
					slog.Int("status", apiErr.HTTPStatus),
					slog.String("code", apiErr.Code),
					slog.String("err_msg", err.Error()),
				)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(apiErr.HTTPStatus)
			_ = json.NewEncoder(w).Encode(apiErr)
		}
		return nil
	}

	wrapped := Handler(errorHandlingWrapper)

	for i := len(b.mws) - 1; i >= 0; i-- {
		wrapped = b.mws[i](wrapped)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		_ = wrapped(w, r)
	}
}
