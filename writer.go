package gluon

import (
	"fmt"
	"strings"
)

type Writer struct {
	sb      *strings.Builder
	indent  int
	newLine bool
}

func NewWriter() *Writer {
	return &Writer{sb: &strings.Builder{}}
}

func (w *Writer) Indent(i int) {
	w.indent += i
}

func (w *Writer) ShiftLeft() {
	w.Indent(-2)
}

func (w *Writer) ShiftRight() {
	w.Indent(2)
}

func (w *Writer) Printf(str string, args ...interface{}) {
	if w.newLine {
		for i := 0; i < w.indent; i++ {
			w.sb.WriteByte(' ')
		}
	}
	w.sb.WriteString(fmt.Sprintf(str, args...))
	w.newLine = strings.HasSuffix(str, "\n")
}

func (w *Writer) String() string {
	return w.sb.String()
}
