package gluon

type Package struct {
	name     string
	parent   *Package
	classes  []*Class
	packages []*Package
}

func NewPackage(name string, parent *Package) *Package {
	p := &Package{name: name, parent: parent}
	if parent != nil {
		parent.packages = append(parent.packages, p)
	}
	return p
}

func (p *Package) toClassDiagram(w *Writer) {
	w.Printf("package \"%s\" #FFFFFF {\n", p.name)
	for _, class := range p.classes {
		class.toClassDiagram(w)
	}
	w.Printf("\n")
}
