package slogctx

import (
	"context"
	"log/slog"
)

type Handle = func(context.Context, slog.Record) error
type HandleConv = func(Handle) Handle

type RecordPrepare = func(context.Context, slog.Record) slog.Record

func PrepareConv(prepare RecordPrepare) HandleConv {
	return func(fn Handle) Handle {
		return func(ctx context.Context, rec slog.Record) error {
			return fn(ctx, prepare(ctx, rec))
		}
	}
}
