package slogctx

import (
	"testing"
)

func TestDefault(t *testing.T) {
	defaultBackup := defaultNamespace
	defer func() {
		defaultNamespace = defaultBackup
	}()

	if Default() != defaultBackup {
		t.Errorf("Default() should return defaultNamespace")
	}

	ns := NewNamespace()
	SetDefault(ns)
	if Default() != ns {
		t.Errorf("Default() should return the namespace set by SetDefault()")
	}
}
