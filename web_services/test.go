package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type foo struct {
	Message string `json: "message,omitempty"` // JSON field names do not have to match Go struct tags
	Age int
	Name string
	Surname string // apparently, it must be "Surname" and not "surname"
	// only fields starting w/caps are exported
}

type Product struct {
    ProductID      int    `json:"productId"`
    Manufacturer   string `json:"manufacturer"`
    Sku            string `json:"sku"`
    Upc            string `json:"upc"`
    PricePerUnit   string `json:"pricePerUnit"`
    QuantityOnHand int    `json:"quantityOnHand"`
    ProductName    string `json:"productName"`
}

func main() {

	data, _ := json.Marshal( &foo{ "4Score", 56, "Abe", "Lincoln" } )
	fmt.Println( "Here is result of json.Marshal: ", string( data ) )
	f := foo{}
	err := json.Unmarshal( []byte( `{"Message": "4Score", "Age": 56, "Name": "Abe"}` ), &f )
	if err != nil {
		log.Fatal( err )
	} else {
		fmt.Println( "Here is json.Unmarshal: " + f.Message )
	}

	fmt.Println()
	
	product := &Product{
		ProductID:      123,
		Manufacturer:   "Acme Corp",
		PricePerUnit:   "12.99",
		Sku:            "451qw456",
		Upc:            "1234569",
		QuantityOnHand: 24,
		ProductName:    "Gizmo",
	}

	productJSON, err := json.Marshal( product )
	if err != nil {
		log.Fatal( err )
	} else {
		fmt.Println( "Here is our product: " + string( productJSON ) )
	}

	productJSONSrc := `{
        "productId":      456,
        "manufacturer":   "Jack Corp",
        "sku":            "451qw457",
        "upc":            "1234568",
        "pricePerUnit":   "14.99",
        "quantityOnHand": 64,
        "productName":    "GizmoGadget"
    }`

	newProduct := Product{}
	// cannot reuse name "err"
	err2 := json.Unmarshal( []byte( productJSONSrc ), &newProduct )
	if err2 != nil {
		log.Fatal( err2 )
	} else {
		fmt.Println( "Product name from JSON: " + newProduct.ProductName  )
		// fmt.Println( "Product from JSON: " + newProduct  )
	}
}

