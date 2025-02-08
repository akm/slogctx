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

// HandleConv is a function that converts a Handle function.
type HandlerConv = func(slog.Handler) slog.Handler

// NewHandlerConv returns a HandlerConv function from a HandleConv function.
func NewHandlerConv(fn HandleConv) HandlerConv {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}
