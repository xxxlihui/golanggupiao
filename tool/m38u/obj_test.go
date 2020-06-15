package main

import "testing"

func TestObj(t *testing.T) {
	a := &AA{}
	a.C = "123"
	println(a.C)
}

type A struct {
	C string
}
type B struct {
	C string
}
type AA struct {
	A
	C string
	B
}
