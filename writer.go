package gluon

import (
	"fmt"
	"strings"
)

type Writer struct {
	sb *strings.Builder
}

func NewWriter() *Writer {
	return &Writer{sb: &strings.Builder{}}
}

func (w *Writer) Printf(str string, args ...interface{}) {
	w.sb.WriteString(fmt.Sprintf(str, args...))
}

func (w *Writer) String() string {
	return w.sb.String()
}
