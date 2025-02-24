package slogctx

import (
	"context"
	"log/slog"
)

// Handle is a function as same as slog.Handler.Handle.
type Handle = func(context.Context, slog.Record) error

// HandleFactory is a function that creates a Handle function with a Handler.
type HandleFactory = func(slog.Handler) Handle

// HandleConv is a function that converts a Handle function.
type HandleConv = func(HandleFactory) HandleFactory

// RecordConv is a function that converts a Record function with [context.Context].
type RecordConv = func(context.Context, slog.Record) slog.Record

// RecordHandleConv converts a RecordConv function to a HandleConv function.
func RecordHandleConv(prepare RecordConv) HandleConv {
	return func(fn HandleFactory) HandleFactory {
		return func(h slog.Handler) Handle {
			return func(ctx context.Context, rec slog.Record) error {
				return fn(h)(ctx, prepare(ctx, rec))
			}
		}
	}
}
