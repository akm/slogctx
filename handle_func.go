package slogw

import (
	"context"
	"log/slog"
)

type HandleFunc = func(context.Context, slog.Record) error
