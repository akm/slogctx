package slogw

import "log/slog"

type factory struct {
	HandlerWrapFuncs
}

func (f *factory) RegisterWrapFunc(fn HandlerWrapFunc) {
	f.HandlerWrapFuncs = append(f.HandlerWrapFuncs, fn)
}

func (f *factory) Register(fn func(HandleFunc) HandleFunc) {
	f.RegisterWrapFunc(newWrapFunc(fn))
}

func newWrapFunc(fn func(orig HandleFunc) HandleFunc) HandlerWrapFunc {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}
