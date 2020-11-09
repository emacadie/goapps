package product

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"shelfunit.info/golang/inventoryservice/cors"
	"strconv"
	"strings"
)

const productsBasePath = "products"

const fileNameService = "product.service"

func SetupRoutes(apiBasePath string) {
	handleProducts := http.HandlerFunc(productsHandler)
	handleProduct  := http.HandlerFunc(productHandler)
	// "products" handles group
	// "products/" handles one
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsBasePath), cors.Middleware(handleProducts))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsBasePath), cors.Middleware(handleProduct))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	var funcName = "In " + fileNameService + ".productHandler: "
	log.Println( funcName )
	urlPathSegments := strings.Split(r.URL.Path, "/products/")
	productID, err  := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])

	if err != nil {
		log.Print( err )
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// return a single product
		log.Println( funcName + "In http.MethodGet" )
		product, err := getProduct( productID ) 
		if product == nil {
			log.Println( "Product is nil" )
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err2 := json.Marshal( product )
		if err2 != nil {
			log.Print( err2 )
			log.Println( " " )
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err2 = w.Write( j )
		if err2 != nil {
			log.Fatal( err2 )
		}

	case http.MethodPut:
		log.Println( funcName + "In http.MethodPut" )
		var product Product
		err :=  json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Printf( "%s Error: %s\n", funcName, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if product.ProductID != productID {
			log.Printf( "%s product.ProductID %d is not equal to productID %d \n", funcName, product.ProductID, productID )
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateProduct(product)
		if err != nil {
			log.Printf( "%s Error: %s\n", funcName, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
	case http.MethodOptions:
		return

	case http.MethodDelete:
		removeProduct(productID)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} // switch

} // productHandler

func productsHandler(w http.ResponseWriter, r *http.Request) {
	var funcName = "In " + fileNameService + ".productsHandler: "
	log.Println( funcName )
	switch r.Method {
	case http.MethodGet:
		productList, errP := getProductList()
		if errP != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	case http.MethodPost:
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var productID = 0
		productID, err = insertProduct(product)
		log.Printf( "%s inserted product %d \n", funcName, productID )
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	} // switch
} // productsHandler

