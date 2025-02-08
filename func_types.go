package slogctx

import (
	"context"
	"log/slog"
)

type Handle = func(context.Context, slog.Record) error
type SlogHandleConv = func(Handle) Handle

type RecordPrepare = func(context.Context, slog.Record) slog.Record

func PrepareConv(prepare RecordPrepare) SlogHandleConv {
	return func(fn Handle) Handle {
		return func(ctx context.Context, rec slog.Record) error {
			return fn(ctx, prepare(ctx, rec))
		}
	}
}

type HandlerConv = func(slog.Handler) slog.Handler

func NewHandlerConv(fn SlogHandleConv) HandlerConv {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}

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
