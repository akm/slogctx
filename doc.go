/*
Package slogctx provides a simple way to log context information with the slog package.

[slogctx.Add] functin is used to add a function to add a value to slog.Record from context.Context.
[slogctx.New] function is used to instantiate slog.Logger with a Handler extended the Handler given as an argument.

See examples for more information.

[slogctx.Add] and [slogctx.New] are package-level functions. [slogctx.Add] affects all logger instances created by [slogctx.New].
If you want to add a function to a specific logger instance, use [slogctx.Namespace]. [slogctx.Namespace] has more primitive
functions such as [Namespace.AddHandlerConv] than [slogctx.Add]. You can create a namespace by [slogctx.NewNamespace].
*/
package slogctx
