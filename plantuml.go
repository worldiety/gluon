package gluon

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type UmlDiagram struct {
	Name   string
	src    string
	buf    []byte
	format string
}

func (d *UmlDiagram) Write(w io.Writer) (int, error) {
	return w.Write(d.buf)
}

func (d *UmlDiagram) WriteFile(f string) (e error) {
	file, err := os.Create(f)
	if err != nil {
		return fmt.Errorf("cannot create %s: %e", f, err)
	}
	defer DeferClose(err, &e)
	_, err = d.Write(file)
	if err != nil {
		return fmt.Errorf("cannot write %s: %e", f, err)
	}
	return nil
}

func (d *UmlDiagram) String() string {
	return d.src
}

func ClassDiagram(p *Package, format string) (*UmlDiagram, error) {
	w := NewWriter()
	w.Printf("@startuml\n")
	style(w)
	p.toClassDiagram(w)
	w.Printf("@enduml\n")
	b, err := PlantUml(w.String(), format)

	return &UmlDiagram{
		Name:   "ClassDiagram",
		src:    w.String(),
		buf:    b,
		format: format,
	}, err
}

func style(w *Writer) {
	w.Printf(`
		skinparam shadowing false
		skinparam class {
			BackgroundColor #6AC4F2
			ArrowColor SeaGreen
			BorderColor #000000
			
		}
`)
}

// PlantUml invokes plantuml to render the src into the given format (png,svg,eps).
func PlantUml(src string, format string) ([]byte, error) {
	cmd := exec.Command("plantuml", "-t"+format, "-p")
	cmd.Env = os.Environ()
	cmd.Stdin = bytes.NewBufferString(src)
	out := bytes.NewBuffer(nil)
	errOut := bytes.NewBuffer(nil)
	cmd.Stdout = out
	cmd.Stderr = errOut

	err := cmd.Run()
	if err != nil {
		tmp := string(out.Bytes()) + string(errOut.Bytes())
		return nil, fmt.Errorf("failed to render '%s' (%s): %e", src, tmp, err)
	}
	return out.Bytes(), nil
}
