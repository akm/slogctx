package slogw

var defaultWrapFuncs HandlerWrapFuncs

func RegisterWrapFunc(f HandlerWrapFunc) {
	defaultWrapFuncs = append(defaultWrapFuncs, f)
}

func Register(f func(HandleFunc) HandleFunc) {
	RegisterWrapFunc(NewWrapFunc(f))
}
