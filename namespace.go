package slogctx

import "log/slog"

type Namespace struct {
	HandlerWrapFuncs
}

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (f *Namespace) RegisterHandlerWrapFunc(fn HandlerWrapFunc) {
	f.HandlerWrapFuncs = append(f.HandlerWrapFuncs, fn)
}

func (f *Namespace) RegisterHandlerPrepareFunc(fn RecordPrepare) {
	f.RegisterHandleFuncWrapFunc(PrepareConv(fn))
}

func (f *Namespace) RegisterHandleFuncWrapFunc(fn SlogHandleConv) {
	f.RegisterHandlerWrapFunc(NewWrapFunc(fn))
}

func (f *Namespace) Register(fn RecordPrepare) {
	f.RegisterHandlerPrepareFunc(fn)
}

func NewWrapFunc(fn SlogHandleConv) HandlerWrapFunc {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}
