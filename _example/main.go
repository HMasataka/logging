package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/HMasataka/logging"
)

type Struct struct {
	Number int64
	String string
}

func main() {
	logger := slog.New(logging.NewHandler(slog.NewJSONHandler(os.Stdout, nil)))
	slog.SetDefault(logger)
	A()
}

func A() {
	ctx := logging.WithValue(context.Background(), "number", 12)
	ctx = logging.WithValue(ctx, "string", "data")
	ctx = logging.WithValue(ctx, "struct", Struct{
		Number: 42,
		String: "struct_data",
	})

	B(ctx)
}

func B(ctx context.Context) {
	slog.ErrorContext(ctx, "this is an error")
	slog.InfoContext(ctx, "this is info")
}
