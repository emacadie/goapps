package main

import(
	"fmt"
	"strings"
)

func main() {
	sampleString := "Never trust a programmer who carries a screwdriver\n"
	fmt.Println( "Before: " + sampleString )
	strLowerCase := strings.ToLower( sampleString )
	fmt.Println( "Lower case: " + strLowerCase )
	strUpperCase := strings.ToUpper( sampleString )
	fmt.Println( "Upper case: " + strUpperCase )
	strTitleCase := strings.Title( sampleString )
	fmt.Println( "Title case: " + strTitleCase )

	sampleString02 := "welcome to the dollhouse\n"
	fmt.Println( "should 'to' and 'the' be capitalized?: " + strings.Title( sampleString02 ) )
	fmt.Println( "Using function properTitle: " + properTitle( sampleString02 + "\n" ) )
}

func properTitle( input string ) string {
	words := strings.Fields( input )
	smallwords := " a an on the to "
	// smallwords := strings.Fields( " a an on the to " )
	for index, word := range words {
		if strings.Contains( smallwords, " "+word+" " ) {
			words[ index ] = word
		} else {
			words[ index ] = strings.Title( word )
		}
	}
	return strings.Join( words, " " )
}


