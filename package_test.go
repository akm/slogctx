package slogctx

import (
	"context"
	"log/slog"
	"testing"
)

func TestPackageFunctions(t *testing.T) {
	t.Parallel()
	backupNamespace := defaultNamespace
	t.Cleanup(func() { defaultNamespace = backupNamespace })
	SetDefault(NewNamespace())
	testAddAndNew(t, Add, New)
}

func TestDefault(t *testing.T) {
	t.Parallel()
	defaultBackup := defaultNamespace
	defer func() {
		defaultNamespace = defaultBackup
	}()

	defaultHandlerConvsLen := len(*defaultNamespace)

	if Default() != defaultBackup {
		t.Errorf("Default() should return defaultNamespace")
	}

	ns := NewNamespace()
	SetDefault(ns)
	if Default() != ns {
		t.Errorf("Default() should return the namespace set by SetDefault()")
	}

	Add(func(ctx context.Context, rec slog.Record) slog.Record { return rec })
	if len(*ns) != 1 {
		t.Errorf("Add() should append a handler converter to the namespace")
	}
	if len(*defaultBackup) != defaultHandlerConvsLen {
		t.Errorf("Add() should not affect the default namespace")
	}
}
