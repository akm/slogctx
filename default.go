package slogw

import "log/slog"

var defaultTransformFuncs WrapFuncs

func RegisterHandleTransformFunc(f WrapFunc) {
	defaultTransformFuncs = append(defaultTransformFuncs, f)
}

func WrapHandler(h slog.Handler) slog.Handler {
	return defaultTransformFuncs.Wrap(h)
}

func Register(f func(HandleFunc) HandleFunc) {
	RegisterHandleTransformFunc(NewHandleWrapFunc(f))
}
