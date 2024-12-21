package slogw

import "log/slog"

var defaultNamespace = &Namespace{}

func Default() *Namespace {
	return defaultNamespace
}

func SetDefault(ns *Namespace) {
	defaultNamespace = ns
}

func RegisterHandlerWrapFunc(f HandlerWrapFunc) {
	defaultNamespace.RegisterHandlerWrapFunc(f)
}

func RegisterHandlerPrepareFunc(f HandlePrepareFunc) {
	defaultNamespace.RegisterHandlerPrepareFunc(f)
}

func RegisterHandleFuncWrapFunc(fn HandleFuncWrapFunc) {
	defaultNamespace.RegisterHandleFuncWrapFunc(fn)
}

func Register(f HandlePrepareFunc) {
	defaultNamespace.Register(f)
}

func Wrap(h slog.Handler) slog.Handler {
	return defaultNamespace.Wrap(h)
}

func New(h slog.Handler) *slog.Logger {
	return defaultNamespace.New(h)
}
