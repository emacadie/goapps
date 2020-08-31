package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "This is my string. There are many strings like it, but this one is mine"
	fmt.Println( "Here is original string: ", sampleString )
	// notice this uses "=", and NOT ":="
	// I originally mistyped ":=", and got this output:
	// no new variables on left side of :=
	sampleString = strings.Replace( sampleString, "string", "compiler", -1 )
	// the -1 means replace all. A positive number would replace that many times
	fmt.Println( "Here is new string: ", sampleString )
}


