package main

import (
	"fmt"
	"log"
	"net/http"
	"shelfunit.info/golang/inventoryservice/database"
	"shelfunit.info/golang/inventoryservice/product"
	"shelfunit.info/golang/inventoryservice/receipt"
	"time"
	// note the underscore
	_"github.com/jackc/pgx/v4" // v4.9.2 // indirect
	// _"github.com/lib/pq" // v1.8.0 // indirect this one likes $1
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
	database.SetupDatabase()

	log.Println( "Got the db" )
	fmt.Println( "test" )
	// register foo handler
	http.Handle( "/foo", &fooHandler{ Message: "Foo called" } )
	// register bar handler
	http.HandleFunc( "/bar", barHandler )

	product.SetupRoutes( apiBasePath )
	receipt.SetupRoutes( apiBasePath )

	http.ListenAndServe( ":5000", nil )
}

