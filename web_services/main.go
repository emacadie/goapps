package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct {
	Message string
}

func ( f *fooHandler ) ServeHTTP( w http.ResponseWriter, r *http.Request ) {
	w.Write( []byte ( f.Message ) )
	
}

func main() {
	fmt.Println( "test" )
	// register handler
	http.Handle( "/foo", &fooHandler{ Message: "Foo called" } )
	// call servmux
	http.ListenAndServe( ":5000", nil )
}

