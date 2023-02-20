package main

import (
	"fmt"
)

type NumType int64

type Describable interface {
	Describe() *NumType
}

func (num *NumType) Describe() *NumType {
	fmt.Printf("NumType type %T - value representation %v\n", num, num)
	fmt.Printf("pointer value %v - type %T\n", *(num), *(num))
	fmt.Printf("num value  %v - type %T\n\n", num, num)
	fmt.Println("-- We are still in the describe function of NumType")
	fmt.Printf("-- Inside, functions can be different as long as they have same interface signature\n\n")
	return num
}

type NumStruct struct {
	num int64
}

func (s *NumStruct) Describe() *NumType {
	fmt.Printf("NumStruct type %T - value representation %v\n", s, s)
	fmt.Printf("pointer value %v - type %T\n", *(s), *(s))
	fmt.Printf("s.num value  %v - type %T\n\n", s.num, s.num)
	_converted := NumType(s.num)
	return &_converted
}

func main() {
	var n1 Describable = &NumStruct{10}
	r1 := n1.Describe()
	fmt.Printf("n1 has returned a pointer to value %v of type %T\n\n", *r1, r1)

	var n2 Describable
	_n2 := NumType(10)
	n2 = &_n2
	r2 := n2.Describe()
	fmt.Printf("n2 has returned a pointer to value %v of type %T\n\n", *r2, r2)

	// ! Panics!
	// InitOnlyIf()
}
