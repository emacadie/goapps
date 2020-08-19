package main

import (
	"fmt"
)

func main() {
	var firstname string
	var lastname string
	fmt.Println( "What is your name?" )
	// this scans until it gets a new line
	// you can use this to format and validate input
	// you can use %q for a quoted string
	// bufio might be better
	fmt.Scanf( "%s %s", &firstname, &lastname )
	fmt.Printf( "Hello, %s %s, nice to meet you. \n", firstname, lastname )
	
	var firstNumber int
	var secondNumber int
	fmt.Println( "What two numbers would you like to add?" )
	fmt.Scanf( "%d %d", &firstNumber, &secondNumber )
	fmt.Printf( "Your total adding %d and %d is: %d\n", firstNumber, secondNumber, ( firstNumber + secondNumber ) )
	
}

