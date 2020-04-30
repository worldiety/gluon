package gluon

import (
	"strings"
)

func lcase(str string) string {
	return strings.ToLower(str)
}

func firstUpper(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(str[0:1]) + str[1:]
}

func join(sep string, elems ...string) string {
	return strings.Join(elems, sep)
}
