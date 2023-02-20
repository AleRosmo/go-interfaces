package main

import "fmt"

// This interface identifies all the types that implements a different method signature,
// that allows for arguments to be passed and returned from the function rather than
// using a custom type.
// ! Doesn't work

type DescribableParam interface {
	DescribeParam(int64) int64
}
type TypeDescribable Describable

// ? I'm not sure what the above does is it simply a Type Definition
// ? of an Interface Type, so simply another interface?
// ? Is this also any useful?
// * Without instantiating a type during interface assignation it panics
// * Arguably very logic since the beginning.

func InitOnlyIf() {

	// Panics, not a valid initialization, the interface is created but
	// a type implementing the interface is not instantiated
	var td TypeDescribable
	td.Describe()

	// Also panics if we initialize it without assigning a type, using only the interface.
	var n3 DescribableParam
	r3 := n3.DescribeParam(10)
	fmt.Printf("n3 has returned a pointer to value %v of type %T\n\n", r3, r3)
}
