package gluon

type Class struct {
	name   string
	parent *Package
	fields []*Field
}

func NewClass(name string, parent *Package) *Class {
	c := &Class{name: name, parent: parent}
	if parent != nil {
		parent.classes = append(parent.classes, c)
	}
	return c
}

func (c *Class) AddFields(fields ...*Field) *Class {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *Class) toClassDiagram(w *Writer) {
	w.Printf("class %s {\n", c.name)
	for _, field := range c.fields {
		field.toClassDiagram(w)
	}
	w.Printf("}\n")
}
