package gluon

import "strings"

type Qualifier []string

func NewQualifier(s ...string) Qualifier {
	return s
}

func (q Qualifier) String() string {
	return strings.Join(q, ".")
}
