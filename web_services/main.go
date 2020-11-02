package main

import (
 	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	"net/http"
	"shelfunit.info/golang/inventoryservice/product"
	// "strconv"
	// "strings"
	"time"
)


/*
var productList []Product

func init() {
	productsJSON := `[  
        {
			"productId": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "p5z343vdS",
			"upc": "939581000000",
			"pricePerUnit": "497.45",
			"quantityOnHand": 9703,
			"productName": "sticky note"
		  },
		  {
			"productId": 2,
			"manufacturer": "Hessel, Schimmel and Feeney",
			"sku": "i7v300kmx",
			"upc": "740979000000",
			"pricePerUnit": "282.29",
			"quantityOnHand": 9217,
			"productName": "leg warmers"
		  },
		  {
			"productId": 3,
			"manufacturer": "Swaniawski, Bartoletti and Bruen",
			"sku": "q0L657ys7",
			"upc": "111730000000",
			"pricePerUnit": "436.26",
			"quantityOnHand": 5905,
			"productName": "lamp shade"
		  }  
        ]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}



*/

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

	// productListHandler := http.HandlerFunc( productsHandler )
	// productItemHandler := http.HandlerFunc( productHandler  )

	// http.Handle( "/products", middlewareHandler( productListHandler ) )
	// http.Handle( "/products/", middlewareHandler( productItemHandler ) )
	// call servmux
	http.ListenAndServe( ":5000", nil )
}

