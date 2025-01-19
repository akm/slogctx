package slogctx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	defaultBackup := defaultNamespace
	defer func() {
		defaultNamespace = defaultBackup
	}()

	assert.Equal(t, defaultBackup, Default())

	ns := NewNamespace()
	SetDefault(ns)
	assert.Equal(t, ns, Default())
}
