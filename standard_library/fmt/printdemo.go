package main

import (
	"fmt"
)

func main() {
	var age = 43
	// print returns the bytes written
	var out, _ = fmt.Print( "Jeremy is ", age, " years old\n" )
	print( "Bytes written: ", out )
	print( "\n" )
}

