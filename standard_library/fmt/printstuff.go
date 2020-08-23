package main

import (
	"fmt"
)

type point struct {
	x, y int
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := point{ 2, 3 }
	fmt.Printf( "Here is our point: %v \n", p )
	// must use braces to instantiate a struct
	newPerson := Person{ "Rob", "Pike", 65 } // just guessing on his age
	fmt.Printf( "Our person: %v is a %T\n", newPerson, newPerson )

	var isCool = true
	fmt.Printf( "Our boolean: %t \n", isCool )

	fmt.Printf( "Here is a number: %d \n", 4567 )
	fmt.Printf( "Same number in binary: %b \n", 4567 )
	fmt.Printf( "Use %s to get ASCII char from a number: %c \n", "%c", 33 )
	fmt.Printf( "Use %s to print hex values: %x \n", "%x", "\n" )
}


