package main

import (
	"context"
	"log/slog"
	"os"
)

var keys = []string{"user_id", "request_id"}

type MyLogHandler struct {
	slog.Handler
}

var _ slog.Handler = &MyLogHandler{}

func (h *MyLogHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, key := range keys {
		if v := ctx.Value(key); v != nil {
			r.AddAttrs(slog.Attr{Key: string(key), Value: slog.AnyValue(v)})
		}
	}

	return h.Handler.Handle(ctx, r)
}

func main() {
	logger := slog.New(&MyLogHandler{slog.NewJSONHandler(os.Stderr, nil)})
	slog.SetDefault(logger)
	A()
}

func A() {
	ctx := context.WithValue(context.Background(), "user_id", 1)
	ctx = context.WithValue(ctx, "request_id", "000")
	B(ctx)
}

func B(ctx context.Context) {
	slog.InfoContext(ctx, "Hello, world!")
}
