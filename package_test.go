package slogctx

import (
	"context"
	"log/slog"
	"testing"
)

func TestPackageFunctions(t *testing.T) { //nolint:paralleltest,tparallel
	// Dn't run this test in parallel.
	// This test use package-level functins Add and New.
	// So, they might be affected by other tests
	// if they are run in parallel.

	backupNamespace := defaultNamespace
	t.Cleanup(func() { defaultNamespace = backupNamespace })
	SetDefault(NewNamespace())
	testAddAndNew(t, Add, New)
}

func TestDefault(t *testing.T) { //nolint:paralleltest
	defaultBackup := defaultNamespace
	t.Cleanup(func() { defaultNamespace = defaultBackup })

	defaultHandlerConvsLen := len(*defaultNamespace)

	if Default() != defaultBackup {
		t.Errorf("Default() should return defaultNamespace")
	}

	ns := NewNamespace()
	SetDefault(ns)
	if Default() != ns {
		t.Errorf("Default() should return the namespace set by SetDefault()")
	}

	Add(func(_ context.Context, rec slog.Record) slog.Record { return rec })
	if len(*ns) != 1 {
		t.Errorf("Add() should append a handler converter to the namespace")
	}
	if len(*defaultBackup) != defaultHandlerConvsLen {
		t.Errorf("Add() should not affect the default namespace")
	}
}
