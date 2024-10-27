package slogw

import "log/slog"

type WrapFunc = func(slog.Handler) slog.Handler

type WrapFuncs []WrapFunc

func (fns WrapFuncs) Wrap(h slog.Handler) slog.Handler {
	for i := len(fns) - 1; i >= 0; i-- {
		h = fns[i](h)
	}
	return h
}
