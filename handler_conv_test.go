package slogctx

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestWrapperWithAttrsAndWithGroup(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	handler0 := slog.NewJSONHandler(buf, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
	wrapper := newWrapper(handler0, func(h slog.Handler) Handle {
		return func(ctx context.Context, r slog.Record) error {
			return h.Handle(ctx, r)
		}
	})
	t.Run("WithAttrs", func(t *testing.T) {
		buf.Reset()
		l0 := slog.New(wrapper)
		l1 := l0.With("key1", "value1")
		l1.Info("test")
		if buf.Len() == 0 {
			t.Error("buf.Len() == 0")
		}
		t.Logf("buf.String() = %v", buf.String())
		res := map[string]interface{}{}
		err := json.Unmarshal(buf.Bytes(), &res)
		if err != nil {
			t.Errorf("json.Unmarshal() failed: %v", err)
		}
		if res["key1"] != "value1" {
			t.Errorf("res[\"key1\"] expected to be \"value1\", but got %v", res["key1"])
		}
	})
	t.Run("WithGroup", func(t *testing.T) {
		buf.Reset()
		l0 := slog.New(wrapper)
		l1 := l0.WithGroup("group1").With("key1", "value1")
		l1.Info("test")
		if buf.Len() == 0 {
			t.Error("buf.Len() == 0")
		}
		t.Logf("buf.String() = %v", buf.String())
		res := map[string]interface{}{}
		err := json.Unmarshal(buf.Bytes(), &res)
		if err != nil {
			t.Errorf("json.Unmarshal() failed: %v", err)
		}
		if res["group1"] == nil {
			t.Fatalf("res[\"group1\"] expected to be not nil, but got nil")
		}
		g, ok := res["group1"].(map[string]interface{})
		if !ok {
			t.Fatalf("res[\"group1\"] expected to be map[string]interface{}, but got %T", res["group1"])
		}
		if g["key1"] != "value1" {
			t.Errorf("g[\"key1\"] expected to be \"value1\", but got %v", g["key1"])
		}
	})
}
