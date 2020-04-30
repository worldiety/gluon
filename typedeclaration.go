package gluon

import "strings"



// TypeDec (type declaration) wraps the full qualified type name and its generic parameters
type TypeDec struct {
	pointer    bool
	name       Qualifier
	parameters []TypeParameter
}

func NewTypeDec(name Qualifier, params ...TypeParameter) TypeDec {
	return TypeDec{
		name:       name,
		parameters: params,
	}
}

func (t TypeDec) Name() Qualifier {
	return t.name
}

func (t TypeDec) Parameters() []TypeParameter {
	return t.parameters
}

func (t TypeDec) String() string {
	sb := &strings.Builder{}
	if t.pointer {
		sb.WriteString("*")
	}
	sb.WriteString(t.name.String())
	if len(t.parameters) > 0 {
		sb.WriteString("<")
		for i, tp := range t.parameters {
			sb.WriteString(tp.String())
			if i < len(t.parameters)-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString(">")
	}
	return sb.String()
}

// TypeParameter reflects the current generics proposal and also describes build-in generics.
//   Proposal for function
//     func func1(type T)(s T)
//     func func2(type T stringer)(s T)
//     type MyType(type T1, type T2 stringer) // T1 and T2 are the names of the used contracts
//
//   build-in
//     func func1 (s []interface{}) // slice is generic and interface{} its contract
//     func func2 (s []stringer) // slice is generic and stringer its contract
//     type MyType map[string]int // string and int are Contracts and name is empty
type TypeParameter struct {
	Name     string
	Contract Qualifier
}

func (t TypeParameter) String() string {
	sb := &strings.Builder{}
	if len(t.Name) > 0 {
		sb.WriteString(t.Name)
		sb.WriteString(" ")
	}
	if len(t.Contract) > 0 {
		sb.WriteString(t.Contract.String())
	}
	return sb.String()
}
