package slogctx

import "log/slog"

type Namespace struct {
	HandlerConvs
}

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (x *Namespace) Add(fn RecordConv) {
	x.AddHandleConv(PrepareConv(fn))
}

func (x *Namespace) AddHandleConv(fn HandleConv) {
	x.AddHandlerConv(NewHandlerConv(fn))
}

func (x *Namespace) AddHandlerConv(fn HandlerConv) {
	x.HandlerConvs = append(x.HandlerConvs, fn)
}

func (x *Namespace) New(h slog.Handler) *slog.Logger {
	return x.HandlerConvs.New(h)
}
