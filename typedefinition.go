package gluon

// A TypeDef (type definition) describes the nature of a type
type TypeDef struct {
	visibility     Visibility
	name           string
	parent         *Package
	fields         []*Field
	underlyingType *TypeDec
}

// NewTypeDefinition defines a new type
func NewTypeDef(visibility Visibility, name string) *TypeDef {
	c := &TypeDef{name: name, visibility: visibility}
	return c
}

func (c *TypeDef) AddFields(fields ...*Field) *TypeDef {
	c.fields = append(c.fields, fields...)
	return c
}

func (c *TypeDef) toClassDiagram(w *Writer) {
	w.Printf("class %s {\n", c.name)
	w.ShiftRight()
	for _, field := range c.fields {
		field.toClassDiagram(w)
	}
	w.ShiftLeft()
	w.Printf("}\n")
}

func (c *TypeDef) Qualifier() Qualifier {
	return append(c.parent.Qualifier(), c.name)
}

