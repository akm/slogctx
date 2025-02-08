package slogctx

import "log/slog"

var defaultNamespace = &Namespace{}

func Default() *Namespace {
	return defaultNamespace
}

func SetDefault(ns *Namespace) {
	defaultNamespace = ns
}

func Add(f RecordPrepare) {
	defaultNamespace.Add(f)
}

func New(h slog.Handler) *slog.Logger {
	return defaultNamespace.New(h)
}
