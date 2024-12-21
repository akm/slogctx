package slogw

import "log/slog"

var defaultFactory = &Namespace{}

func RegisterHandlerWrapFunc(f HandlerWrapFunc) {
	defaultFactory.RegisterHandlerWrapFunc(f)
}

func RegisterHandlerPrepareFunc(f HandlePrepareFunc) {
	defaultFactory.RegisterHandlerPrepareFunc(f)
}

func RegisterHandleFuncWrapFunc(fn HandleFuncWrapFunc) {
	defaultFactory.RegisterHandleFuncWrapFunc(fn)
}

func Register(f HandlePrepareFunc) {
	defaultFactory.Register(f)
}

func Wrap(h slog.Handler) slog.Handler {
	return defaultFactory.Wrap(h)
}

func New(h slog.Handler) *slog.Logger {
	return slog.New(Wrap(h))
}
