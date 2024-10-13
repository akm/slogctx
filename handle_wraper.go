package slogwrap

import (
	"context"
	"log/slog"
)

type HandleFunc = func(context.Context, slog.Record) error

type handleWrapper struct {
	slog.Handler
	handle HandleFunc
}

var _ slog.Handler = (*handleWrapper)(nil)

func NewHandleTransformFunc(fn func(orig HandleFunc) HandleFunc) TransformFunc {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &handleWrapper{Handler: h, handle: handle}
	}
}

func (h *handleWrapper) Handle(ctx context.Context, rec slog.Record) error {
	return h.handle(ctx, rec)
}
