package slogctx_test

import (
	"context"
	"log/slog"
	"os"

	"github.com/akm/slogctx"
)

// An example of logging with a value in context.
func Example() {
	// Define a empty anonymous struct as key for value in context.
	type ctxKeyType struct{}
	ctxKey := ctxKeyType{}

	// Add a function to add a value to slog.Record from context.Context.
	slogctx.Add(func(ctx context.Context, rec slog.Record) slog.Record {
		val, ok := ctx.Value(ctxKey).(string)
		if ok {
			rec.Add("key1", val)
		}
		return rec
	})
	// TextHandler with testOptions have a function to remove "time" from log.
	hander := slog.NewTextHandler(os.Stdout, testOptions)

	// Instantiate slog.Logger by using slog.New.
	logger0 := slog.New(hander)
	// Instantiate slog.Logger by using slogctx.New.
	logger1 := slogctx.New(hander)

	// Log with a value in context.
	ctx1 := context.WithValue(context.TODO(), ctxKey, "value1")
	logger0.InfoContext(ctx1, "blah blah")
	logger1.InfoContext(ctx1, "blah blah")

	// Output:
	// level=INFO msg="blah blah"
	// level=INFO msg="blah blah" key1=value1
}
