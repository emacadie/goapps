package main

// to run: go run template.go 

import (
	"fmt"
	"os"
	"html/template"
)

type BlogPost struct {
	Title string
	Content string
}

func main() {
	post := BlogPost{ "First Post", "This is the post content" }
	fmt.Println( "Hello, playground" )
	tmpl, err := template.New( "blog-tmpl" ).Parse(`<h1>{{.Title}}</h1><div><p>{{.Content}}</p></div>`)
	if err != nil {
		panic( err )
	}
	err = tmpl.Execute( os.Stdout, post )
	if err != nil {
		panic( err )
	}
	fmt.Println( " " )
}

