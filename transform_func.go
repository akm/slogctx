package slogwrap

import "log/slog"

type TransformFunc = func(slog.Handler) slog.Handler

type TransformFuncs []TransformFunc

func (fns TransformFuncs) Wrap(h slog.Handler) slog.Handler {
	for i := len(fns) - 1; i >= 0; i-- {
		h = fns[i](h)
	}
	return h
}
