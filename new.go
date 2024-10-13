package slogwrap

import "log/slog"

func New(h slog.Handler) *slog.Logger {
	return slog.New(WrapHandler(h))
}
