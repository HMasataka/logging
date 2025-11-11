package logging

import (
	"context"
	"log/slog"
	"sync"
)

func NewHandler(handler slog.Handler) slog.Handler {
	return &LogHandler{
		Handler: handler,
	}
}

type LogHandler struct {
	slog.Handler
}

var _ slog.Handler = &LogHandler{}

func (h *LogHandler) Handle(ctx context.Context, r slog.Record) error {
	if v, ok := ctx.Value(loggingKey).(*sync.Map); ok {
		v.Range(func(key, val any) bool {
			if keyString, ok := key.(string); ok {
				r.AddAttrs(slog.Any(keyString, val))
			}

			return true
		})
	}

	return h.Handler.Handle(ctx, r)
}
