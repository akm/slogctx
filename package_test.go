package slogctx

import (
	"context"
	"log/slog"
	"testing"
)

func TestDefault(t *testing.T) {
	defaultBackup := defaultNamespace
	defer func() {
		defaultNamespace = defaultBackup
	}()

	defaultHandlerConvsLen := len(defaultNamespace.HandlerConvs)

	if Default() != defaultBackup {
		t.Errorf("Default() should return defaultNamespace")
	}

	ns := NewNamespace()
	SetDefault(ns)
	if Default() != ns {
		t.Errorf("Default() should return the namespace set by SetDefault()")
	}

	Add(func(ctx context.Context, rec slog.Record) slog.Record { return rec })
	if len(ns.HandlerConvs) != 1 {
		t.Errorf("Add() should append a handler converter to the namespace")
	}
	if len(defaultBackup.HandlerConvs) != defaultHandlerConvsLen {
		t.Errorf("Add() should not affect the default namespace")
	}
}
