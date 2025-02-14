# slogctx

[![CI](https://github.com/akm/slogctx/actions/workflows/ci.yml/badge.svg)](https://github.com/akm/slogctx/actions/workflows/ci.yml)
[![codecov](https://codecov.io/github/akm/slogctx/graph/badge.svg?token=9BcanbSLut)](https://codecov.io/github/akm/slogctx)
[![Enabled Linters](https://img.shields.io/badge/dynamic/yaml?url=https%3A%2F%2Fraw.githubusercontent.com%2Fakm%2Fslogctx%2Frefs%2Fheads%2Fmain%2F.project.yaml&query=%24.linters&label=enabled%20linters&color=%2317AFC2)](.golangci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/akm/slogctx)](https://goreportcard.com/report/github.com/akm/slogctx)
[![Documentation](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/akm/slogctx)
[![license](https://img.shields.io/github/license/akm/slogctx)](./LICENSE)

`slogctx` means `slog and context`. It provides a simple way to log context information with the slog package.

## Install

```
go get github.com/akm/slogctx@latest
```

## Usage

You can register your preparing function like this:

```golang
	slogctx.Add(func(ctx context.Context, rec slog.Record) slog.Record {
		val, ok := ctx.Value(ctxKey1).(string)
		if ok {
			rec.Add("key1", val)
		}
		return rec
	})
```

And you can get a logger working with your handle function by using `slogctx.New` instead of `slog.New` .

```golang
    handler := slog.NewTextHandler(writer, nil)
    logger := slogctx.New(handler)
```

`writer` must be a io.Writer like os.Stdout, bytes.Buffer or etc.

```golang
	ctx := context.WithValue(context.Background(), ctxKey1, "value1")
	logger.InfoContext(ctx, "foo")
```

Then `logger` outputs a log with `foo` as `msg` and `value1` as `key1` like this:

```
time=2024-12-21T12:18:51.893+09:00 level=INFO msg=foo key1=value1
```

See [example_test.go](./example_test.go) for more detail.

## Alternatives

- https://github.com/PumpkinSeed/slog-context
  - Log all data in context.Context without specifying key

## CONTRIBUTING

If you find a bug, typo, incorrect test, have an idea, or want to help with an existing issue, please create an issue or pull request.

## LICENSE

[MIT](./LICENSE)
