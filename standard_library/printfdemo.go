package main

import (
	"fmt"
)

func main() {
	var age = 42
	var name = "joe"
	// Printf can infer the types w/%v
	fmt.Printf( "My name is %v and my age is %v \n", name, age )
	// we can also use %s for string and %d for integer
	// reversing it gives this: y name is %!d(string=joe) and my age is %!s(int=42)
	fmt.Printf( "My name is %s and my age is %d \n", name, age )
	// you should really check the inputs before printing
	// or give them types
	var age2 int = 10
	var name2 string = "John"
	fmt.Printf( "My name is %s and my age is %d \n", name2, age2 )

	// making pi 22/7 gives 3
	var pi float32 =  3.141592
	fmt.Printf( "Pi is %f \n", pi )
	// to show 2 digits:
	fmt.Printf( "Pi to 2 digits is %.2f \n", pi )
	fmt.Println( "ints in a table" )
	fmt.Printf( "|%d|%d|%d|\n", 23, 44, 59 )
	fmt.Println( "variable width floats, Go pads w/zeros" )
	fmt.Printf( "|%f|%f|%f|\n", 23.1234, 442.36, 5964.2 )
	fmt.Println( "let's try 2 rows with varying widths" )
	fmt.Printf( "|%f|%f|%f|\n", 23.1234, 442.36, 5964.2 )
	fmt.Printf( "|%f|%f|%f|\n", 98.999, 12.3456, 64.02 )
	fmt.Println( "use Printf to get it all right" )
	fmt.Printf( "|%4.2f|%4.2f|%4.2f|\n", 23.1234, 442.36, 5964.2 )
	fmt.Printf( "|%4.2f|%4.2f|%4.2f|\n", 98.999, 12.3456, 64.02 )
	fmt.Println( "still not lined up, so try 7.2" )
	fmt.Printf( "|%7.2f|%7.2f|%7.2f|\n", 23.1234, 442.36, 5964.2 )
	fmt.Printf( "|%7.2f|%7.2f|%7.2f|\n", 98.999, 12.3456, 64.02 )
	fmt.Println( "-7.2 for left justified" )
	fmt.Printf( "|%-7.2f|%-7.2f|%-7.2f|\n", 23.1234, 442.36, 5964.2 )
	fmt.Printf( "|%-7.2f|%-7.2f|%-7.2f|\n", 98.999, 12.3456, 64.02 )
	fmt.Println( "This works with strings" )
	fmt.Printf( "|%-7s|%-7s|%-7s|\n", "qqq", "www", "eee" )
	fmt.Printf( "|%-7s|%-7s|%-7s|\n", "r", "tt", "yyyy" )
	fmt.Println( "change s to q for quoted string" )
	fmt.Printf( "|%-7s|%-7s|%-7s|\n", "qqq", "www", "eee" )
	fmt.Printf( "|%-7s|%-7s|%-7q|\n", "r", "tt", "yyyy" )

	output := fmt.Sprintf( "|%-7s|%-7s|%-7s|\n", "qqq", "www", "eee" )
	fmt.Println( "Here is output: ", output )
}


