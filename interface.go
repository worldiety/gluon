package gluon

type Interface struct {
	name    string
	methods []*InterfaceMethod
}

func NewInterface(name string) *Interface {
	return &Interface{name: name}
}

func (i *Interface) WriteGo(w *Writer) {
	w.Printf("type %s interface {\n", firstUpper(i.name))
	w.ShiftRight()
	for _, m := range i.methods {
		m.WriteGo(w)
	}
	w.ShiftLeft()
	w.Printf("}\n")
}

func (i *Interface) AddMethods(methods ...*InterfaceMethod) *Interface {
	for _, m := range methods {
		i.methods = append(i.methods, m)
		m.parent = i
	}
	return i
}
