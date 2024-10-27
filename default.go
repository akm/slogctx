package slogw

var defaultWrapFuncs WrapFuncs

func RegisterWrapFunc(f WrapFunc) {
	defaultWrapFuncs = append(defaultWrapFuncs, f)
}

func Register(f func(HandleFunc) HandleFunc) {
	RegisterWrapFunc(NewWrapFunc(f))
}
