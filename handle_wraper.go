package slogw

import (
	"context"
	"log/slog"
)

type HandleFunc = func(context.Context, slog.Record) error

type wrapper struct {
	slog.Handler
	handle HandleFunc
}

var _ slog.Handler = (*wrapper)(nil)

func NewWrapFunc(fn func(orig HandleFunc) HandleFunc) WrapFunc {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}

func (h *wrapper) Handle(ctx context.Context, rec slog.Record) error {
	return h.handle(ctx, rec)
}
