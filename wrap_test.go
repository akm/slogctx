package slogwrap

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapWithRegister(t *testing.T) {
	logging := func(ctx context.Context) {
		slog.InfoContext(ctx, "test")
	}

	type ctxKey1Tyep struct{}
	ctxKey1 := ctxKey1Tyep{}

	addKeyValue := func(ctx context.Context, v string) context.Context {
		return context.WithValue(ctx, ctxKey1, v)
	}

	defaultTransformFuncs = nil
	Register(
		func(orig HandleFunc) HandleFunc {
			return func(ctx context.Context, rec slog.Record) error {
				val, ok := ctx.Value(ctxKey1).(string)
				if ok {
					rec.Add("key1", val)
				}
				return nil
			}
		},
	)

	newLoggerAndBuf := func() (*slog.Logger, *bytes.Buffer) {
		buf := bytes.NewBufferString("")
		logger := New(slog.NewJSONHandler(buf, nil))
		return logger, buf
	}

	t.Run("no value in context", func(t *testing.T) {
		logger, buf := newLoggerAndBuf()
		slog.SetDefault(logger)
		ctx := context.Background()
		logging(ctx)
		d := map[string]any{}
		err := json.Unmarshal(buf.Bytes(), &d)
		assert.NoError(t, err)
		d["key1"] = nil
	})

	t.Run("value in context", func(t *testing.T) {
		logger, buf := newLoggerAndBuf()
		slog.SetDefault(logger)
		baseCtx := context.Background()
		t.Run("value1", func(t *testing.T) {
			ctx := addKeyValue(baseCtx, "value1")
			logging(ctx)
			d := map[string]any{}
			err := json.Unmarshal(buf.Bytes(), &d)
			assert.NoError(t, err)
			assert.Equal(t, "value1", d["key1"])
		})
		t.Run("value2", func(t *testing.T) {
			ctx := addKeyValue(baseCtx, "value2")
			logging(ctx)
			d := map[string]any{}
			err := json.Unmarshal(buf.Bytes(), &d)
			assert.NoError(t, err)
			assert.Equal(t, "value2", d["key1"])
		})
	})

}
