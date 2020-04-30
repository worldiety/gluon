package gluon

type InterfaceMethod struct {
	name   string
	doc    *Documentation
	parent *Interface
}

func NewInterfaceMethod(name string) *InterfaceMethod {
	return &InterfaceMethod{name: name}
}

func (m *InterfaceMethod) WriteGo(w *Writer) {
	w.Printf("func %s(", m.name)
	w.Printf(")\n")
}
