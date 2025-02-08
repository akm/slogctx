package slogctx

type Namespace struct {
	HandlerConvs
}

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (f *Namespace) RegisterHandlerWrapFunc(fn HandlerConv) {
	f.HandlerConvs = append(f.HandlerConvs, fn)
}

func (f *Namespace) AddHandleConv(fn HandleConv) {
	f.RegisterHandlerWrapFunc(NewHandlerConv(fn))
}

func (f *Namespace) Register(fn RecordPrepare) {
	f.AddHandleConv(PrepareConv(fn))
}
