package slogw

import "log/slog"

var defaultFactory = &Factory{}

func RegisterHandlerWrapFunc(f HandlerWrapFunc) {
	defaultFactory.RegisterHandlerWrapFunc(f)
}

func RegisterHandlerPrepareFunc(f HandlePrepareFunc) {
	defaultFactory.RegisterHandlerPrepareFunc(f)
}

func Register(f HandleFuncWrapFunc) {
	defaultFactory.Register(f)
}

func Wrap(h slog.Handler) slog.Handler {
	return defaultFactory.Wrap(h)
}

func New(h slog.Handler) *slog.Logger {
	return slog.New(Wrap(h))
}
