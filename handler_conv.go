package slogctx

import (
	"context"
	"log/slog"
)

type wrapper struct {
	impl          slog.Handler
	handle        Handle
	handleFactory func(slog.Handler) Handle
}

var _ slog.Handler = (*wrapper)(nil)

func newWrapper(h slog.Handler, f func(slog.Handler) Handle) *wrapper {
	return &wrapper{
		impl:          h,
		handleFactory: f,
		handle:        f(h),
	}
}

func (h *wrapper) Enabled(ctx context.Context, lv slog.Level) bool {
	return h.impl.Enabled(ctx, lv)
}

func (h *wrapper) Handle(ctx context.Context, rec slog.Record) error {
	return h.handle(ctx, rec)
}

func (h *wrapper) WithAttrs(attrs []slog.Attr) slog.Handler {
	return newWrapper(h.impl.WithAttrs(attrs), h.handleFactory)
}

func (h *wrapper) WithGroup(name string) slog.Handler {
	return newWrapper(h.impl.WithGroup(name), h.handleFactory)
}

// HandleConv is a function that converts a Handle function.
type HandlerConv = func(slog.Handler) slog.Handler

// NewHandlerConv returns a HandlerConv function from a HandleConv function.
func NewHandlerConv(fn HandleConv) HandlerConv {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{impl: h, handle: handle}
	}
}
