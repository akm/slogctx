package slogw

import "log/slog"

func Wrap(h slog.Handler) slog.Handler {
	return defaultWrapFuncs.Wrap(h)
}

func New(h slog.Handler) *slog.Logger {
	return slog.New(Wrap(h))
}
