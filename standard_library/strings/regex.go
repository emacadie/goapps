package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleString := "This is my string. There are many strings like it, but this one is mine"
	// reg ex for string starting w/s, has letters, ends with g
	r, _ := regexp.Compile( `s([a-z]+)g` )
	fmt.Println( "Does string match regex?: ", r.MatchString( sampleString ) )
	fmt.Println( r.FindAllString( sampleString, -1 ) )


	sampleString02 := "This is my song. There are many strings like it, but this one is mine"
	fmt.Println( "Changed first 'string' in string to 'song'. Now trying regex" )
	fmt.Println( r.FindAllString( sampleString02, -1 ) )
	
	// start w/capital T, ends with r
	r02, _ := regexp.Compile( `T([a-z]+)r` )
	fmt.Println( "New criteria: start w/capital T, ends with r" )
	fmt.Println( r02.FindAllString( sampleString, -1 ) )
	fmt.Println( "So far, we have looked at character sequences" )
	
	fmt.Println( "Now reg ex for word starting w/s, has letters, ends with g at word boundry" )
	r03, _ := regexp.Compile( `s(\w[a-z]+)g\b` )
	fmt.Println( "Using second string" )
	fmt.Println( r03.FindAllString( sampleString02, -1 ) )
	fmt.Println( "Use regexp.FindStringIndex to get the index of string" )
	fmt.Println( r03.FindStringIndex( sampleString02 ) )

	fmt.Println( "ReplaceAllString replaces string" )
	newText := r03.ReplaceAllString( sampleString, "laptop" )
	fmt.Println( "New string: ", newText )
}


