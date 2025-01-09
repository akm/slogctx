package slogw

import "log/slog"

var defaultNamespace = &Namespace{}

func Default() *Namespace {
	return defaultNamespace
}

func SetDefault(ns *Namespace) {
	defaultNamespace = ns
}

func Register(f HandlePrepareFunc) {
	defaultNamespace.Register(f)
}

func New(h slog.Handler) *slog.Logger {
	return defaultNamespace.New(h)
}
