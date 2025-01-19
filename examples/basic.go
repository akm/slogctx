package main

import (
	"context"
	"log/slog"
	"os"

	slogctx "github.com/akm/slogw"
)

func basic() {
	type ctxKey1Type struct{}
	var ctxKey1 = ctxKey1Type{}

	slogctx.Register(func(ctx context.Context, rec slog.Record) slog.Record {
		val, ok := ctx.Value(ctxKey1).(string)
		if ok {
			rec.Add("key1", val)
		}
		return rec
	})

	// Set the logger instantiated by slogctx.New as the default logger.
	slog.SetDefault(slogctx.New(slog.NewTextHandler(os.Stdout, nil)))

	ctx0 := context.Background()
	ctx1 := context.WithValue(ctx0, ctxKey1, "value1")
	slog.InfoContext(ctx1, "foo")
	slog.InfoContext(ctx1, "bar")

	ctx2 := context.WithValue(ctx0, ctxKey1, "value2")
	slog.InfoContext(ctx2, "baz")
}
