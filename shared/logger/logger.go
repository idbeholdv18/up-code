package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/idbeholdv18/up-code/shared/httpkit"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID := httpkit.GetRequestID(ctx); reqID != "" {
		r.AddAttrs(slog.String("request_id", reqID))
	}
	return h.Handler.Handle(ctx, r)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{Handler: h.Handler.WithGroup(name)}
}

func Init(serviceName string) {
	baseHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	ctxHandler := &ContextHandler{Handler: baseHandler}
	log := slog.New(ctxHandler).With(slog.String("service", serviceName))
	slog.SetDefault(log)
}
