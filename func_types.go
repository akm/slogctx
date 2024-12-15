package slogw

import (
	"context"
	"log/slog"
)

type HandleFunc = func(context.Context, slog.Record) error

type HandlerWrapFunc = func(slog.Handler) slog.Handler

type HandlerWrapFuncs []HandlerWrapFunc

func (fns HandlerWrapFuncs) Wrap(h slog.Handler) slog.Handler {
	for i := len(fns) - 1; i >= 0; i-- {
		h = fns[i](h)
	}
	return h
}
