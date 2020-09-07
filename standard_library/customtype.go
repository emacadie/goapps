package main

import (
	"fmt"
	"media"
)

func main() {
	fmt.Println( "My favorite movie" )
	// old way, before we had the Catalogable interface
	// myFave := media.NewMovie( "Farewell My Concubine", media.R, 43.2 )

	// his files have "var", but he did not need it on the screen
	var myFave media.Catalogable = &media.Movie{}
	myFave.NewMovie( "Farewell My Concubine", media.R, 43.2 )

	/* // the fields were mutable before
	myFave.Title = "Farewell My Concubine"
	myFave.Rating = media.R
	myFave.BoxOffice = 43.2
	fmt.Printf( "My favorite movie is: %s\n", myFave.Title )
	fmt.Printf( "It was rated %v\n", myFave.Rating )
	fmt.Printf( "It made %f at the box office\n", myFave.BoxOffice )
    */

	fmt.Printf( "My favorite movie is: %s\n", myFave.GetTitle() )
	fmt.Printf( "It was rated %v\n", myFave.GetRating() )
	fmt.Printf( "It made %f at the box office\n", myFave.GetBoxOffice() )


	// try a different movie
	myFave.SetTitle( "Austin Powers: The Spy Who Shagged Me" )
	myFave.SetRating( media.PG13 )
	myFave.SetBoxOffice( 21.3 )
	fmt.Printf( "My favorite movie is: %s\n", myFave.GetTitle() )
	fmt.Printf( "It was rated %v\n", myFave.GetRating() )
	fmt.Printf( "It made %f at the box office\n", myFave.GetBoxOffice() )


}


