package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	sampleString := "I really like the Go language. It's one of my favorites"
	/*
	searchTerm := "Go"
	result := strings.Contains( sampleString, searchTerm )
	fmt.Println( "Here is sampleString: ", sampleString )
	fmt.Println( "Here is searchTerm: ", searchTerm )
	fmt.Printf( "The sample text includes '%s': %t \n ", searchTerm, result )
*/
	// now he changes it (same file; I hate that) to use a command line arg
	if len( os.Args ) > 1 {
		searchTerm := os.Args[ 1 ]
		result := strings.Contains( sampleString, searchTerm )
		if result {
			fmt.Printf( "The sample text includes '%s': result is %t \n", searchTerm, result )
		} else {
			fmt.Printf( "The sample text does not include '%s': result is %t \n", searchTerm, result )
		}
		prefixResult := strings.HasPrefix( sampleString, searchTerm )
		if prefixResult {
			fmt.Printf( "The sample text starts with '%s': result is %t \n", searchTerm, prefixResult )
		} else {
			fmt.Printf( "The sample text does start with '%s': result is %t \n", searchTerm, prefixResult )
		}

		suffixResult := strings.HasSuffix( sampleString, searchTerm )
		if prefixResult {
			fmt.Printf( "The sample text ends with '%s': result is %t \n", searchTerm, suffixResult )
		} else {
			fmt.Printf( "The sample text does end with '%s': result is %t \n", searchTerm, suffixResult )
		}

	} else {
		fmt.Println( "You must enter a search term" )
	}
}

