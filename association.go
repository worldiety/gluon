package gluon

type Multiplicity struct{
	Min int
	Max int
}


type Association interface {
}

// A Composition is a 0..n association between two classes where exactly one parent has the ownership.
type Composition struct {
	Owner *TypeDef
	OwnerMultiplicity Multiplicity

}

type bla struct{
	X string
}

type blub struct{
	X string
}

type Xer1 interface{
	X() string
}

type Xer2 interface{
	X() string
}

type IBla interface{
	Do( struct{X string})
}

type IBlub interface{
	Do( struct{X string})
}

func a(){
	a := bla{}
	a = bla(blub{})
	_ = a

	var i IBla
	var i2 IBlub

	i=i2
	i.Do(bla{})
	i.Do(blub{})
}