package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "this is a string"
	string2 := "this is another string"
	/*
	if string1 == string2 {
		fmt.Println( "The strings are identical" )
	} else {
		fmt.Println( "The strings do not match" )
	}
*/
	areStringsIdentical( string1, string2 )
	areStringsIdentical( "this is a string", "this is a string" )
	// it's case sensitive
	areStringsIdentical( "this is a string", "This is a string" )
	
	// strings.Compare can be faster than using comparison operators
	stooges := []string{ "Larry", "Curly", "Moe" }
	for _, stooge := range stooges {
		fmt.Println( "Comparing 'Larry' to '" + stooge + "' gives: ", strings.Compare( "Larry", stooge ) )
	}

	// sometimes you want case-insensitive comparison
	fmt.Println( CompareCaseIns( "Hey, this is a string", "I like turtles" ) )
	fmt.Println( CompareCaseIns( "four", "five" ) )
	fmt.Println( CompareCaseIns( "I LIKE GO", "I like Go" ) )
	
}

func CompareCaseIns( a, b string ) bool {
	fmt.Printf( "In CompareCaseIns with '%s' and '%s': ", a, b )
	if len( a ) != len ( b ) {
		return false
	} else if strings.ToLower( a ) == strings.ToLower( b ) {
		return true
	} else {
		return false
	}
}

func areStringsIdentical( stringA string, stringB string) {
	fmt.Printf( "Comparing \"%s\" and \"%s\": ", stringA, stringB )
	// if stringA == stringB {
	// using Compare
	// "The result will be 0 if a==b, -1 if a < b, and +1 if a > b."
	// I really hate these functions
	if strings.Compare( stringA, stringB ) == 0 {
		fmt.Println( "The strings are identical" )
	} else {
		fmt.Println( "The strings do not match" )
	}
}


