package slogctx

import "log/slog"

// Namespace is a slice of HandlerConv.
type Namespace []HandlerConv

// NewNamespace returns a new Namespace.
func NewNamespace() *Namespace {
	return &Namespace{}
}

// AddRecordConv appends a RecordConv function to the Namespace.
func (x *Namespace) AddRecordConv(fn RecordConv) {
	x.AddHandleConv(RecordHandleConv(fn))
}

// AddHandleConv appends a HandleConv function to the Namespace.
func (x *Namespace) AddHandleConv(fn HandleConv) {
	x.AddHandlerConv(NewHandlerConv(fn))
}

// AddHandlerConv appends a HandlerConv function to the Namespace.
func (x *Namespace) AddHandlerConv(fn HandlerConv) {
	*x = append(*x, fn)
}

// New returns a new Logger with the Namespace.
func (x *Namespace) New(h slog.Handler) *slog.Logger {
	return slog.New(x.Wrap(h))
}

// Wrap returns a new Handler with the Namespace.
func (x *Namespace) Wrap(h slog.Handler) slog.Handler {
	s := *x
	for i := len(s) - 1; i >= 0; i-- {
		h = s[i](h)
	}
	return h
}
