package main

import (
	"fmt"
	"net/http"
	"shelfunit.info/golang/inventoryservice/product"
	"time"
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

func middlewareHandler( handler http.Handler ) http.Handler {
	return http.HandlerFunc( func( w http.ResponseWriter, r *http.Request )  {
		fmt.Println( "Before handler; middleware starting" )
		start := time.Now()
		handler.ServeHTTP( w, r )
		fmt.Printf( "Middleware finished: %s \n", time.Since( start ) )
	})
}

const apiBasePath = "/api"

func main() {
	fmt.Println( "test" )
	// register foo handler
	http.Handle( "/foo", &fooHandler{ Message: "Foo called" } )
	// register bar handler
	http.HandleFunc( "/bar", barHandler )

	product.SetupRoutes( apiBasePath )

	http.ListenAndServe( ":5000", nil )
}

