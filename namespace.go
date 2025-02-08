package slogctx

import "log/slog"

type Namespace struct {
	HandlerConvs
}

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (f *Namespace) Add(fn RecordPrepare) {
	f.AddHandleConv(PrepareConv(fn))
}

func (f *Namespace) AddHandleConv(fn HandleConv) {
	f.AddHandlerConv(NewHandlerConv(fn))
}

func (f *Namespace) AddHandlerConv(fn HandlerConv) {
	f.HandlerConvs = append(f.HandlerConvs, fn)
}

func (f *Namespace) Wrap(h slog.Handler) slog.Handler {
	return f.HandlerConvs.Wrap(h)
}

func (f *Namespace) New(h slog.Handler) *slog.Logger {
	return f.HandlerConvs.New(h)
}
