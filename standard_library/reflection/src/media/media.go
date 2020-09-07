package media

import (
	"strings"
)


type Catalogable interface {
	NewMovie( title string, rating Rating, boxOffice float32 )
	GetTitle() string
	GetRating() string
	GetBoxOffice() float32
	SetTitle( newTitle string )
	SetRating( newRating Rating )
	SetBoxOffice( newBoxOffice float32 )
}

type Movie struct {
	title string // mutable by other files if starting w/capital letter
	rating Rating
	boxOffice float32
}
type Rating string

const (
	R = "R (Restricted)"
	G = "G (General audiences)"
	PG = "PG (Parental Guidance)"
	PG13 = "PG-13 (Parental Caution)"
	NC17 = "NC-17 (No children under 17)"
)

func ( m *Movie ) NewMovie(argTitle string, rating Rating, boxOffice float32) {
	m.title = argTitle
	m.rating = rating
	m.boxOffice = boxOffice
}

/*
old NewMovie before Catalogable interface
func NewMovie(argTitle string, rating Rating, boxOffice float32) Movie {
	return Movie {
		title: argTitle,
		rating: rating,
		boxOffice: boxOffice,
	}
}
*/
func ( m *Movie ) GetTitle() string {
	return strings.Title( m.title )
}
func ( m *Movie ) GetRating() string {
	return string( m.rating )
}
func ( m *Movie ) GetBoxOffice() float32 {
	return m.boxOffice
}
// to get the changes to stick, use a pointer-based receiver
// we do need to change them for the setters to see changes,
// but for consistency it is good to change the getters too
// we also could have changed the methods to return a new immutable object
func ( m *Movie ) SetTitle( newTitle string ) {
	m.title = newTitle
}

func ( m *Movie ) SetRating( newRating Rating ) {
	m.rating = newRating
}

func ( m *Movie ) SetBoxOffice( newBoxOffice float32 ) {
	m.boxOffice = newBoxOffice
}


