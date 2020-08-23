package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println( "What is your name" )
	// this scans until it gets a new line
	// you can use this to format and validate input
	// you can use %q for a quoted string
	// bufio might be better
	inputs, _ := fmt.Scanf( "%s", &name )
	switch inputs {
		case 0:
		    fmt.Println( "You must enter a name" )
		case 1:
		    fmt.Printf( "Hello, %s, nice to meet you. You input %d args\n", name, inputs )
	}
	
}

