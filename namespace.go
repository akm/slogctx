package slogctx

type Namespace struct {
	HandlerConvs
}

func NewNamespace() *Namespace {
	return &Namespace{}
}

func (f *Namespace) AddHandlerConv(fn HandlerConv) {
	f.HandlerConvs = append(f.HandlerConvs, fn)
}

func (f *Namespace) AddHandleConv(fn HandleConv) {
	f.AddHandlerConv(NewHandlerConv(fn))
}

func (f *Namespace) Register(fn RecordPrepare) {
	f.AddHandleConv(PrepareConv(fn))
}
