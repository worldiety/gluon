package gluon

type Visibility int

const Public Visibility = 0
const PackagePrivate Visibility = 1

type Field struct {
	v    Visibility
	name string
	dec  TypeDec
}

func NewField(visibility Visibility, name string, dec TypeDec) *Field {
	return &Field{v: visibility, name: name, dec: dec}
}

func (f *Field) toClassDiagram(w *Writer) {
	if f.v == Public {
		w.Printf("+")
	} else {
		w.Printf("~")
	}
	w.Printf("%s : %s\n", f.name, f.dec.String())
}
