package slogctx

import "log/slog"

var defaultNamespace = &Namespace{}

// Default returns the default Namespace.
func Default() *Namespace {
	return defaultNamespace
}

// SetDefault sets the default Namespace.
func SetDefault(ns *Namespace) {
	defaultNamespace = ns
}

// Add appends a RecordConv function to the default Namespace.
func Add(f RecordConv) {
	defaultNamespace.AddRecordConv(f)
}

// New returns a new Logger with the default Namespace.
func New(h slog.Handler) *slog.Logger {
	return defaultNamespace.New(h)
}
