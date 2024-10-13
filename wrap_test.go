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
				return orig(ctx, rec)
			}
		},
	)

	newLoggerAndBuf := func() (*slog.Logger, *bytes.Buffer) {
		buf := bytes.NewBufferString("")
		logger := New(slog.NewJSONHandler(buf, nil))
		return logger, buf
	}

	type pattern struct {
		ctx       context.Context
		key1Value *string
		name      string
	}

	strPtr := func(s string) *string {
		return &s
	}

	patterns := []pattern{
		{context.Background(), nil, "no value"},
		{addKeyValue(context.Background(), "value1"), strPtr("value1"), "value1"},
		{addKeyValue(context.Background(), "value2"), strPtr("value2"), "value2"},
	}
	for _, ptn := range patterns {
		t.Run(ptn.name, func(t *testing.T) {
			logger, buf := newLoggerAndBuf()
			slog.SetDefault(logger)
			logging(ptn.ctx)
			d := map[string]any{}
			t.Logf("buf: %s\n", buf.String())
			err := json.Unmarshal(buf.Bytes(), &d)
			assert.NoError(t, err)
			if ptn.key1Value == nil {
				assert.Nil(t, d["key1"])
			} else {
				assert.Equal(t, *ptn.key1Value, d["key1"])
			}
		})
	}
}
