package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// args[0] is the name of the app, but we don't need that
	// so make a slice for everything from os.Args[ 1 ] onward.
	args := os.Args[ 1: ]
	if ( len( args ) == 1 ) && ( args[ 0 ] == "/help" ) {
		fmt.Println( "Usage: dinnertotal <Total Meal Amount> <Tip Percentage>" );
		fmt.Println( "Example: dinnertotal 20 15" )
	} else 	if len( args ) != 2 {
		fmt.Println( "You must enter 2 arguments." )
		fmt.Println( "Type \"dinnertotal /help\" for more info." )
	} else {
		// args are re-numbered since we put them in a new slice from os.Args
		// args are brought in as strings, so we must convert to floats
		mealTotal, _ := strconv.ParseFloat( args[ 0 ], 32 )
		tipAmount, _ := strconv.ParseFloat( args[ 1 ], 32 )
		fmt.Printf( "Your meal total will be $%.2f\n", calculateTotal( float32( mealTotal ), float32( tipAmount ) ) )
	}
	
}

func calculateTotal( mealTotal float32, tipAmount float32 ) float32 {
	totalPrice := mealTotal + ( mealTotal * ( tipAmount / 100 ) )
	return totalPrice
}


