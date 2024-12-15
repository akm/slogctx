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

func (f *Factory) RegisterHandlerPrepareFunc(fn HandlePrepareFunc) {
	f.RegisterHandlerWrapFunc(f.newWrapFunc(Prepare(fn)))
}

func (f *Factory) RegisterHandleFuncWrapFunc(fn HandleFuncWrapFunc) {
	f.RegisterHandlerWrapFunc(f.newWrapFunc(fn))
}

func (f *Factory) Register(fn HandlePrepareFunc) {
	f.RegisterHandlerPrepareFunc(fn)
}

func (*Factory) newWrapFunc(fn HandleFuncWrapFunc) HandlerWrapFunc {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}
