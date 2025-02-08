package slogctx

import "log/slog"

type Namespace []HandlerConv

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (x *Namespace) AddRecordConv(fn RecordConv) {
	x.AddHandleConv(RecordHandleConv(fn))
}

func (x *Namespace) AddHandleConv(fn HandleConv) {
	x.AddHandlerConv(NewHandlerConv(fn))
}

func (x *Namespace) AddHandlerConv(fn HandlerConv) {
	*x = append(*x, fn)
}

func (x *Namespace) New(h slog.Handler) *slog.Logger {
	return slog.New(x.Wrap(h))
}

func (x *Namespace) Wrap(h slog.Handler) slog.Handler {
	s := *x
	for i := len(s) - 1; i >= 0; i-- {
		h = s[i](h)
	}
	return h
}
