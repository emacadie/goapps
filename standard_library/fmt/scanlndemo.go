package main

import (
	"fmt"
)

func main() {
	fmt.Print( "What is your name? (Enter first name and last name) " )
	var firstName string
	var lastName string
	fmt.Scanln( &firstName, &lastName )
	fmt.Printf( "hello, %s %s, nice to meet you\n", firstName, lastName )
}

