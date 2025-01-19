package main

import (
	"context"
	"log/slog"
	"os"

	slogctx "github.com/akm/slogw"
)

func namespace() {
	type ctxKey1Type struct{}
	var ctxKey1 = ctxKey1Type{}

	ns1 := slogctx.NewNamespace()
	ns1.Register(func(ctx context.Context, rec slog.Record) slog.Record {
		val, ok := ctx.Value(ctxKey1).(string)
		if ok {
			rec.Add("key1", val)
		}
		return rec
	})

	logger := ns1.New(slog.NewTextHandler(os.Stdout, nil))

	ctx0 := context.Background()
	ctx1 := context.WithValue(ctx0, ctxKey1, "value3")
	logger.InfoContext(ctx1, "foo")

	ctx2 := context.WithValue(ctx0, ctxKey1, "value4")
	logger.InfoContext(ctx1, "bar")
	logger.InfoContext(ctx2, "baz")
}
