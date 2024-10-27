package slogw

import "log/slog"

var defaultWrapFuncs WrapFuncs

func RegisterWrapFunc(f WrapFunc) {
	defaultWrapFuncs = append(defaultWrapFuncs, f)
}

func WrapHandler(h slog.Handler) slog.Handler {
	return defaultWrapFuncs.Wrap(h)
}

func Register(f func(HandleFunc) HandleFunc) {
	RegisterWrapFunc(NewWrapFunc(f))
}
