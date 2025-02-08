package slogctx

import (
	"context"
	"log/slog"
)

// Handle is a function as same as slog.Handler.Handle.
type Handle = func(context.Context, slog.Record) error

// HandleConv is a function that converts a Handle function.
type HandleConv = func(Handle) Handle

// RecordConv is a function that converts a Record function with [context.Context].
type RecordConv = func(context.Context, slog.Record) slog.Record

// RecordHandleConv converts a RecordConv function to a HandleConv function.
func RecordHandleConv(prepare RecordConv) HandleConv {
	return func(fn Handle) Handle {
		return func(ctx context.Context, rec slog.Record) error {
			return fn(ctx, prepare(ctx, rec))
		}
	}
}
