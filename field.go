package gluon

type Field struct {
	private bool // means package private
	name    string
}

func NewField(private bool, name string) *Field {
	return &Field{private: private, name: name}
}

func (f *Field) toClassDiagram(w *Writer) {
	if f.private {
		w.Printf("~")
	} else {
		w.Printf("+")
	}
	w.Printf("%s\n", f.name)
}
