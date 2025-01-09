# slogw

![CI](https://github.com/akm/slogw/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/github/akm/slogw/graph/badge.svg?token=8KA0MWBH0F)](https://codecov.io/github/akm/slogw)
![license](https://img.shields.io/github/license/akm/slogw)

`slogw` means `slog wrapper`. It supports to wrap the Handle method of slog.Handler interface.
`slogw` is pronounced `slog-wuh` .

## Install

```
go get github.com/akm/slogw@latest
```

## Usage

You can register your handle function like this:

```golang
	slogw.Register(func(ctx context.Context, rec slog.Record) slog.Record {
		val, ok := ctx.Value(ctxKey1).(string)
		if ok {
			rec.Add("key1", val)
		}
		return rec
	})
```

And you can get a logger working with your handle function by using `slogw.New` instead of `slog.New` .

```golang
    yourHandler := slog.NewTextHandler(writer, nil)
    yourLogger := slogw.New(yourHandler)
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

See [examples/basic.go](./examples/basic.go) and [examlpes/namespace.go](./examples/namespace.go) for more detail.
Try `go run ./examples` for `go run ./examples/namespace`
