package slogwrap

import "log/slog"

var defaultTransformFuncs TransformFuncs

func RegisterHandleTransformFunc(f TransformFunc) {
	defaultTransformFuncs = append(defaultTransformFuncs, f)
}

func WrapHandler(h slog.Handler) slog.Handler {
	return defaultTransformFuncs.Wrap(h)
}

func Register(f func(HandleFunc) HandleFunc) {
	RegisterHandleTransformFunc(NewHandleTransformFunc(f))
}
