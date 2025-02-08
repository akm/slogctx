package slogctx

import "log/slog"

type Namespace struct {
	HandlerConvs
}

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (f *Namespace) RegisterHandlerWrapFunc(fn HandlerConv) {
	f.HandlerConvs = append(f.HandlerConvs, fn)
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

func NewWrapFunc(fn SlogHandleConv) HandlerConv {
	return func(h slog.Handler) slog.Handler {
		handle := fn(h.Handle)
		return &wrapper{Handler: h, handle: handle}
	}
}
