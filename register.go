package slogw

var defaultFactory = &factory{}

func RegisterWrapFunc(f HandlerWrapFunc) {
	defaultFactory.RegisterWrapFunc(f)
}

func Register(f func(HandleFunc) HandleFunc) {
	defaultFactory.Register(f)
}
