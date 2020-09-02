package main

import (
	"fmt"
	"strings"
)
func main() {
	sampleString := "     This is our text     "
	fmt.Printf( "This is our string at the beginning: %q\n", sampleString )
	newString := strings.TrimSpace( sampleString )
	fmt.Printf( "This is our string after calling strings.TrimSpace: %q\n", newString )
	newString02 := strings.TrimLeft( sampleString, " " )
	fmt.Printf( "This is our string after calling strings.TrimLeft: %q\n", newString02 )
	newString03 := strings.TrimRight( sampleString, " " )
	fmt.Printf( "This is our string after calling strings.TrimRight: %q\n", newString03 )

	sampleString04 := "https://www.golang.org"
	domainName := strings.TrimPrefix( sampleString04, "https://" )
	fmt.Println( "SampleString04: ", sampleString04, ", domain after calling strings.TrimPrefix: ", domainName )

	sampleString05 := "trim.go"
	fileName := strings.TrimSuffix( sampleString05, ".go" )
	fmt.Println( "SampleString05: ", sampleString05, ", fileName after calling strings.TrimSuffix: ", fileName )
}

