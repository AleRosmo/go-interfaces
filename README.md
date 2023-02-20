# Interfaces - Explanation

##### Self note: Operands in expressions are the elementary values like function, literals.

## Function Signature's anatomy

Given the example:
```golang
func (i myInt) AddInts(i1 int, i2 int) int {return i1 + i2}
```
The '(i myInt)' part is named receiver type argument.

## Methods Sets
A Methods Set is what determines all the methods that can be called on operands of that type, declared by methods with that Type as receiver argument Methods declared on a defined type T are Method Sets on Defined Types T, like '(i myInt)', Methods declared on a pointer of defined type T are Method Sets on Defined Types T, like '(i myInt)'.

## Type Definition 
Defining a new Type NumType allows zero cost conversion between NumType and int64  ? other number types also?
and allows to define new methods - with Type Aliases you can't.
```golang
type NumType int64
```

## Interface Definition 
We are grouping all the types that implements, the method Describe() with this precise signature, so it is said to 'Implement the Describable interface'. When a new type is defined, it must also implement the method Describe() with it's type receiver argument. We can use `any` or `interface{}`.

```golang
type Describable interface {
    Describe() *NumType
}
```

NumType receiver method Describe(), implementation of this method means it will implement the Describable interface.

```golang
func (num *NumType) Describe() *NumType 
    fmt.Printf("NumType type %T - value representation %v\n", num, num)
    fmt.Printf("pointer value %v - type %T\n", *(num), *(num))
    fmt.Printf("num value  %v - type %T\n\n", num, num)
    fmt.Println("-- We are still in the describe function of NumType")
    fmt.Printf("-- Inside, functions can be different as long as they have same interface signature\n\n")
    return num
}
```
A simple Type Definition containing 'num' property of type 'int64'. We'll implement the Describable interface below by implementing the Describe() method.
```golang
type NumStruct struct {
    num int64
}
```

Implementing Describable interface on NumStruct. The Receiver Type is different from the Interface's Describe() method signature therefore we will convert it's type so that *NumStruct type can still implement the Describable interface.

This method is still implemented by the Describable Interface. Even if it's receiver is another type *NumStruct and contains different declarations, statements and/or expressions as long as the signature is
identical to the one implemented by the interface, the type implements the interface, in this case: `Describe() *NumType.`

```golang
func (s *NumStruct) Describe() *NumType {
    fmt.Printf("NumStruct type %T - value representation %v\n", s, s)
    fmt.Printf("pointer value %v - type %T\n", *(s), *(s))
	fmt.Printf("s.num value  %v - type %T\n\n", s.num, s.num)
	_converted := NumType(s.num)
	return &_converted
}
```

Declaring 'n1' var of Describable's interface type. Will first instantiate a new 'NumStruct' type which implements "Describe()" method implemented by the 'Calculable' interface. The type has a 'num' property which
is of type 'int64'. It then assign the reference the memory pointer address of the newly instantiated 'NumStruct'. We call the 'Describe()' method, which will describe the type just instanced.

```golang
func main() {
    var n1 Describable = &NumStruct{10}
    r1 := n1.Describe()
    fmt.Printf("n1 has returned a pointer to value %v of type %T\n\n", *r1, r1)
}
```

Declaring 'n2' var of type 'Describable' interface Declaring and assigning '_n2' as 'NumType' type and value's literal '10' Assigning to 'n2' a reference to '_n2' memory address Calling 'Describe()' method, which will return a converted

```golang
var n2 Describable
_n2 := NumType(10)
n2 = &_n2
r2 := n2.Describe()
fmt.Printf("n2 has returned a pointer to value %v of type %T\n\n", *r2, r2)
} //Closing the main() function.
```
