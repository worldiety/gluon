package gluon

type Package struct {
	name     string
	parent   *Package
	classes  []*TypeDef
	packages []*Package
}

func NewPackage(name string) *Package {
	p := &Package{name: name}
	return p
}

func (p *Package) toClassDiagram(w *Writer) {
	w.Printf("package \"%s\" #FFFFFF {\n", p.name)
	w.ShiftRight()
	for _, class := range p.classes {
		class.toClassDiagram(w)
	}
	for _, pkg := range p.packages {
		pkg.toClassDiagram(w)
	}
	w.ShiftLeft()
	w.Printf("}\n")
}

func (p *Package) Qualifier() Qualifier {
	if p == nil {
		return nil
	}
	return append(p.parent.Qualifier(), p.name)
}

func (p *Package) AddPackages(pkgs ...*Package) *Package {
	for _, pkg := range pkgs {
		p.packages = append(p.packages, pkg)
		pkg.parent = p
	}
	return p
}

func (p *Package) AddTypes(types ...*TypeDef) *Package {
	for _, def := range types {
		p.classes = append(p.classes, def)
		def.parent = p
	}
	return p
}
