package slogw

var defaultFactory = &factory{}

func RegisterHandlerWrapFunc(f HandlerWrapFunc) {
	defaultFactory.RegisterHandlerWrapFunc(f)
}

func Register(f func(HandleFunc) HandleFunc) {
	defaultFactory.Register(f)
}
