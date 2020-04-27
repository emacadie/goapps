package main

import (
	"fmt"
)

type Book struct {
	ID int
	Title string
	Author string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n" +
		"Author:\t\t%q\n" +
		"Published:\t\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book {
	Book {
		ID: 1,
		Title: "The Hitchhikers Guide To something",
		Author: "Some Dude",
		YearPublished: 1979,
	}, 
	Book {
		ID: 2,
		Title: "Effective Java",
		Author: "Josh Bloch",
		YearPublished: 2019,
	},
	Book {
		ID: 3,
		Title: "Go And Go Some More",
		Author: "Rob Pike",
		YearPublished: 2015,
	},
	Book {
		ID: 4,
		Title: "Typing Is Hard",
		Author: "Every Developer",
		YearPublished: 1979,
	},
	Book {
		ID: 5,
		Title: "JSON For Dummies",
		Author: "Former XMLDude",
		YearPublished: 1979,
	},
	Book {
		ID: 6,
		Title: "Book Six",
		Author: "Man Six",
		YearPublished: 1979,
	},
	Book {
		ID: 7,
		Title: "Book Seven",
		Author: "Seven of Nine",
		YearPublished: 1979,
	},
	Book {
		ID: 8,
		Title: "Eight Book",
		Author: "Person Eight",
		YearPublished: 1979,
	},
	Book {
		ID: 9,
		Title: "Nine Very Imaginative",
		Author: "Seven of Nine",
		YearPublished: 1979,
	},
	Book {
		ID: 10,
		Title: "Totally A Ten",
		Author: "Big Ten",
		YearPublished: 1979,
	},
}

