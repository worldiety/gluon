package gluon

var BuildIn = buildInTypes{
	String: NewTypeDec(Qualifier{"string"}),
}

type buildInTypes struct {
	String TypeDec
}
