package slogctx

import (
	"context"
	"log/slog"
)

type wrapper struct {
	slog.Handler
	handle Handle
}

var _ slog.Handler = (*wrapper)(nil)

func (h *wrapper) Handle(ctx context.Context, rec slog.Record) error {
	return h.handle(ctx, rec)
}

type HandlerConv = func(slog.Handler) slog.Handler

func NewHandlerConv(fn HandleConv) HandlerConv {
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
