package slogw

import "log/slog"

func New(h slog.Handler) *slog.Logger {
	return slog.New(Wrap(h))
}
