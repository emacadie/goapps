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
	items := []string{"one - ", "two - ", "three - "}

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

	fmt.Println( "\n-------\nNext template\n" )
	tmpl02S := "{{range .}}{{.}}{{end}}"
	
	tmpl02, _ := template.New( "tmplt" ).Parse( tmpl02S )
	err = tmpl02.Execute( os.Stdout, items )
	fmt.Println( "\n " )
	
	fmt.Println( " " )

	fmt.Println( "\n-------\nNext template\n" )

	tmpl03S := "{{range $index, $element := .}}{{if mod $index 2}}{{.}}{{end}}{{end}}"	
	fm := template.FuncMap{ "mod": func(i, j int) bool { return i % j == 0 } }
	tmpl03, err := template.New( "tmplt" ).Funcs( fm ).Parse( tmpl03S )
	fmt.Printf( "err: %v \n", err )
	fmt.Printf( "Items: %v \n", items )
	fmt.Printf( "tmpl03: %v \n", tmpl03 )
	err03 := tmpl03.Execute( os.Stdout, items )
	fmt.Println( "\n " )
	fmt.Print( err03 )
	fmt.Println( "\n " )

} // end func main

