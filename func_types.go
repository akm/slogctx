package slogctx

import (
	"context"
	"log/slog"
)

type SlogHandle = func(context.Context, slog.Record) error
type SlogHandleConv = func(SlogHandle) SlogHandle

type RecordPrepare = func(context.Context, slog.Record) slog.Record

func PrepareConv(prepare RecordPrepare) SlogHandleConv {
	return func(fn SlogHandle) SlogHandle {
		return func(ctx context.Context, rec slog.Record) error {
			return fn(ctx, prepare(ctx, rec))
		}
	}
}

type HandlerConv = func(slog.Handler) slog.Handler

type HandlerConvs []HandlerConv

func (fns HandlerConvs) Wrap(h slog.Handler) slog.Handler {
	for i := len(fns) - 1; i >= 0; i-- {
		h = fns[i](h)
	}
	return h
}

func (fns HandlerConvs) New(h slog.Handler) *slog.Logger {
	return slog.New(fns.Wrap(h))
}
