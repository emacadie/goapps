package product

import (
	"shelfunit.info/golang/inventoryservice/cors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/*
func getNextID() int {
	highestID := 1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func findProductByID( productID int ) ( *Product, int ) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}
	return nil, 0
}
*/
const productsBasePath = "products"
func SetupRoutes( apiBasePath string ) {
	handleProducts := http.HandlerFunc( productsHandler )
	handleProduct  := http.HandlerFunc( productHandler )
	http.Handle( fmt.Sprintf( "%s/%s", apiBasePath, productsBasePath ), cors.Middleware( handleProducts ) )
	http.Handle( fmt.Sprintf( "%s/%s/", apiBasePath, productsBasePath ), cors.Middleware( handleProduct ) )
}

func productHandler( w http.ResponseWriter, r *http.Request ) {
	urlPathSegments := strings.Split( r.URL.Path, "/products/" )
	productID, err  := strconv.Atoi( urlPathSegments[ len( urlPathSegments ) - 1 ] )
	
	if err != nil {
		w.WriteHeader( http.StatusNotFound )
		return
	}
	product := getProduct( productID )
	if product == nil {
		w.WriteHeader( http.StatusNotFound )
		return
	}
	switch r.Method {
		case http.MethodGet:
			// return a single product
			productJSON, err := json.Marshal( product )
			if err != nil {
				w.WriteHeader( http.StatusInternalServerError )
				return
			}
			w.Header().Set( "Content-Type", "application/json" )
			w.Write( productJSON )
		case http.MethodPut:
		    // update a product
			var updatedProduct Product
			bodyBytes, err := ioutil.ReadAll( r.Body )
			if err != nil {
				w.WriteHeader( http.StatusBadRequest )
				return
			}
			err = json.Unmarshal( bodyBytes, &updatedProduct )
			if err != nil {
				log.Print( err )
				w.WriteHeader( http.StatusBadRequest )
				return
			}
			if updatedProduct.ProductID != productID {
				w.WriteHeader( http.StatusBadRequest )
				return
			}
		addOrUpdateProduct( updatedProduct )

			w.WriteHeader( http.StatusOK )
			return

	case http.MethodOptions:
		return

		case http.MethodDelete:
		removeProduct( productID )

		default:
			w.WriteHeader( http.StatusMethodNotAllowed )
			return
	} // switch
	
} // productHandler

func productsHandler( w http.ResponseWriter, r *http.Request ) {
	switch r.Method {
		case http.MethodGet:
		productList := getProductList()
		    productsJson, err := json.Marshal( productList )
			if err != nil {
				w.WriteHeader( http.StatusInternalServerError )
				return
			}
			w.Header().Set( "Content-Type", "application/json" )
			w.Write( productsJson )
		    case http.MethodPost:
        var product Product
        err := json.NewDecoder(r.Body).Decode(&product)
        if err != nil {
            log.Print(err)
            w.WriteHeader(http.StatusBadRequest)
            return
        } 
        _, err = addOrUpdateProduct(product) 
        if err != nil {
            log.Print(err)
            w.WriteHeader(http.StatusBadRequest) 
            return 
        }
        w.WriteHeader(http.StatusCreated)  
/*
		case http.MethodPost:
			// add a new product to the list
			var newProduct Product
			bodyBytes, err := ioutil.ReadAll( r.Body )
			if err != nil {
				w.WriteHeader( http.StatusBadRequest )
				return
			}
			err = json.Unmarshal( bodyBytes, &newProduct )
			if err != nil {
				w.WriteHeader( http.StatusBadRequest )
				return
			}
			if newProduct.ProductID != 0 {
				w.WriteHeader( http.StatusBadRequest )
				return
			}
		_, err := addOrUpdateProduct( newProduct )
		if err != nil {
			w.WriteHeader( http.StatusBadRequest )
			return
		}
			// newProduct.ProductID = getNextID()
			// productList = append( productList, newProduct )
			w.WriteHeader( http.StatusCreated )
			return
*/
		case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	} // switch
} // productsHandler
