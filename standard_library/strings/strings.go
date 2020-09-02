package main

import (
	"fmt"
)

func main() {
	ourString := "\x47\x6f\x20\x69\x73\x20\x41\x77\x65\x73\x6f\x6d\x65\x21"
	fmt.Println( "ourString is: ", ourString )
	for i := 0; i < len( ourString ); i++ {
		fmt.Printf( "%x ", ourString[ i ] )
	}
	fmt.Println()
	// use %q to get the quoted byte sequence - this can debug non-printable characters
	stringFromGoBlog := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println( "String from Go blog: ", stringFromGoBlog )
	for i := 0; i < len( stringFromGoBlog ); i++ {
		fmt.Printf( "%q ", stringFromGoBlog[ i ] )
	}
	fmt.Println()
	// strings byte values are indexed
	newString := "Strings in go"
	fmt.Println( "Here is newString: ", newString )
	fmt.Println( "Here is newString[ 3 ]: ", newString[ 3 ] )
	// use slice notation to get the characters
	fmt.Println( "Here is newString[ 0:5 ]: ", newString[ 0:5 ] )
	
}


