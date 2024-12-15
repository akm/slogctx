package slogw

import (
	"context"
	"log/slog"
)

type wrapper struct {
	slog.Handler
	handle HandleFunc
}

var _ slog.Handler = (*wrapper)(nil)

func (h *wrapper) Handle(ctx context.Context, rec slog.Record) error {
	return h.handle(ctx, rec)
}
