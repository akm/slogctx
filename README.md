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
	slogw.Register(
		func(orig slogw.HandleFunc) slogw.HandleFunc {
			return func(ctx context.Context, rec slog.Record) error {
				val, ok := ctx.Value(ctxKey1).(string)
				if ok {
					rec.Add("key1", val)
				}
				return orig(ctx, rec)
			}
		},
	)
```

And you can get a logger with your handle function.

```golang
    yourHandler := slog.NewJSONHandler(writer, nil)
    logger := slogw.New(yourHandler)
```

`writer` must be a io.Writer like os.Stdout, bytes.Buffer or etc.
