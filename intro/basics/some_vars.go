package main

import (
	"fmt"
)

func main() {
	var i int
	i = 42
	fmt.Println("i is: ", i)
	
	// float32 or float64
	var f float32 = 3.14
	fmt.Println("f is ", f)
	// you may not need to specify data type
	// implicit initialization syntax
	firstName := "George"
	fmt.Println("Here is firstName: ", firstName)

	// boolean
	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println("complex number: ", c)
	r, i2 := real(c), imag(c)
	fmt.Println("real part of c: ", r, ", imaginary part of c:", i2)

	// pointers
	var firstNameP *string = new(string)
	fmt.Println("here is firstNameP: ", firstNameP)
	// to assign, dereference
	*firstNameP = "George"
	fmt.Println("here is firstNameP: ", firstNameP, ", dereferenced: ", *firstNameP)
	// address of operator:
	firstNameA := "john"
	ptr := &firstNameA
	fmt.Println("ptr: ", ptr, ", derefed: ", *ptr)
	firstNameA = "Thomas"
	fmt.Println("ptr: ", ptr, ", derefed: ", *ptr)

	// constants must be initialized when declared
	// and must be evaluated at compile time
	// so it cannot be the return value of a function
	const pi = 3.1415
	fmt.Println("pi is: ", pi)
	const c2 = 3 
	// implicitly typed constant
	fmt.Println("c2 + 3:", c2 + 3)
	fmt.Println("c2 + 1.2:", c2 + 1.2)
	const c3 int = 3
	fmt.Println("c3 + 3: ", c3 + 3)
	fmt.Println("we convert c3 + 1.2:", float32(c3) + 1.2)
	
	// iota
	
	
}


