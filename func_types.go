package slogctx

import (
	"context"
	"log/slog"
)

type SlogHandle = func(context.Context, slog.Record) error
type SlogHandleConv = func(SlogHandle) SlogHandle

type HandlePrepareFunc = func(context.Context, slog.Record) slog.Record

func Prepare(prepare HandlePrepareFunc) SlogHandleConv {
	return func(fn SlogHandle) SlogHandle {
		return func(ctx context.Context, rec slog.Record) error {
			return fn(ctx, prepare(ctx, rec))
		}
	}
}

type HandlerWrapFunc = func(slog.Handler) slog.Handler

type HandlerWrapFuncs []HandlerWrapFunc

func (fns HandlerWrapFuncs) Wrap(h slog.Handler) slog.Handler {
	for i := len(fns) - 1; i >= 0; i-- {
		h = fns[i](h)
	}
	return h
}

func (fns HandlerWrapFuncs) New(h slog.Handler) *slog.Logger {
	return slog.New(fns.Wrap(h))
}
