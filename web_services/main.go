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

func barHandler( w http.ResponseWriter , r *http.Request ) {
	w.Write( []byte ( "bar called" ) )	
}

func main() {
	fmt.Println( "test" )
	// register foo handler
	http.Handle( "/foo", &fooHandler{ Message: "Foo called" } )
	// register bar handler
	http.HandleFunc( "/bar", barHandler )
	// call servmux
	http.ListenAndServe( ":5000", nil )
}

