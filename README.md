# slogw

`slogw` means `slog wrapper`. It supports to wrap the Handle method of slog.Handler interface.
`slogw` is pronounced `slog-wuh` .

## Install

```
go get github.com/akm/slogw@latest
```

## Usage

You can register your handle function like this:

```golang
	slogw.Register(func(ctx context.Context, rec slog.Record) *slog.Record {
		val, ok := ctx.Value(ctxKey1).(string)
		if ok {
			rec.Add("key1", val)
		}
		return &rec
	})
```

And you can get a logger with your handle function.

```golang
    yourHandler := slog.NewJSONHandler(writer, nil)
    yourLogger := slogw.New(yourHandler)
```

`writer` must be a io.Writer like os.Stdout, bytes.Buffer or etc.

Then yourLogger's logs contain the value named `key1` which is got from ctx by using `ctxKey1`.
