package slogw

import "log/slog"

type Factory struct {
	HandlerWrapFuncs
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) RegisterHandlerWrapFunc(fn HandlerWrapFunc) {
	f.HandlerWrapFuncs = append(f.HandlerWrapFuncs, fn)
}

func (f *Factory) Register(fn HandleFuncWrapFunc) {
	f.RegisterHandlerWrapFunc(newWrapFunc(fn))
}

func newWrapFunc(fn HandleFuncWrapFunc) HandlerWrapFunc {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}
