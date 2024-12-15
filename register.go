package slogw

import "log/slog"

var defaultWrapFuncs HandlerWrapFuncs

func RegisterWrapFunc(f HandlerWrapFunc) {
	defaultWrapFuncs = append(defaultWrapFuncs, f)
}

func Register(f func(HandleFunc) HandleFunc) {
	RegisterWrapFunc(NewWrapFunc(f))
}

func NewWrapFunc(fn func(orig HandleFunc) HandleFunc) HandlerWrapFunc {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}
