package main

import (
	"fmt"
	"github.com/ee4g/macro"
)

//go:generate go run gen.go

func main(){
	macro.MustApply()
	fmt.Println("done")
}