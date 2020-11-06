package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"database/sql"
	"strconv"
	"sync"

	"shelfunit.info/golang/inventoryservice/database"
)

const fileNameData = "In product.data."

var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{m: make(map[int]Product)}

func init() {
	fmt.Println("loading products..")
	prodMap, err := loadProductMap()
	productMap.m = prodMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d products loaded \n", len(productMap.m))
}

func loadProductMap() (map[int]Product, error) {
	fileName2 := "products.json"
	_, err := os.Stat(fileName2)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file[%s] does not exist", fileName2)
	}
	file, _ := ioutil.ReadFile(fileName2)
	productList := make([]Product, 0)
	err = json.Unmarshal([]byte(file), &productList)
	if err != nil {
		log.Fatal(err)
	}
	prodMap := make(map[int]Product)
	for i := 0; i < len(productList); i++ {
		prodMap[productList[i].ProductID] = productList[i]
	}
	return prodMap, nil
}

func getProduct(productID int) (*Product, error) {
	var funcName = fileNameData + "getProduct: "
	log.Printf( funcName + " with product ID %d \n", productID )

	row := database.DbConn.QueryRow( 
		`select productId, manufacturer, sku, upc, pricePerUnit, quantityOnHand, productName 
         from products where productId = $1`, productID )
	product := &Product{}
	err := row.Scan(&product.ProductID, 
			&product.Manufacturer, 
			&product.Sku, 
			&product.Upc, 
			&product.PricePerUnit, 
			&product.QuantityOnHand, 
			&product.ProductName)

	if err == sql.ErrNoRows {
		log.Println( "No rows" )
		return nil, nil
	} else if err != nil {
		log.Fatal( err )
		return nil, err
	}
	return product, nil

	/*
	productMap.RLock()
	defer productMap.RUnlock()
	if product, ok := productMap.m[productID]; ok {
		return &product
	}
	return nil
	*/
} // getProduct

func removeProduct(productID int) {
	var funcName = fileNameData + "removeProduct: "
	log.Printf( "In %s with productID: %d \n", funcName, productID )
	productMap.Lock()
	defer productMap.Unlock()
	delete(productMap.m, productID)
}

func getProductList() ([]Product, error) {
	var funcName = fileNameData + "getProductList: "
	log.Println( funcName )
	results, err := database.DbConn.Query( `select productId, manufacturer, sku, upc, pricePerUnit, quantityOnHand, productName from products` )
	if err != nil {
		return nil, err
	}
	defer results.Close()
	products := make([]Product, 0)
	for results.Next() {
		var product Product
		fmt.Println( "Got a product" )
		results.Scan(&product.ProductID, 
			&product.Manufacturer, 
			&product.Sku, 
			&product.Upc, 
			&product.PricePerUnit, 
			&product.QuantityOnHand, 
			&product.ProductName)
		products = append( products, product )
		fmt.Println( "Appended a product to product list: ", strconv.Itoa( product.ProductID ) )
	}
	return products, nil

	/*
	productMap.RLock()
	products := make([]Product, 0, len(productMap.m))
	for _, value := range productMap.m {
		products = append(products, value)
	}
	productMap.RUnlock() // why not defer? is that only needed if there is a possible error?
	return products
	*/
} // getProductList

func getProductIds() []int {
	productMap.RLock()
	productIds := []int{}
	for key := range productMap.m {
		productIds = append(productIds, key)
	}
	productMap.RUnlock()
	sort.Ints(productIds)
	return productIds
}

func getNextProductID() int {
	productIDs := getProductIds()
	return productIDs[len(productIDs)-1] + 1
}

func addOrUpdateProduct(product Product) (int, error) {
	var funcName = fileNameData + "addOrUpdateProduct: "
	log.Println( funcName )
	// if the product id is set, replace it, otherwise return error
	addOrUpdateID := -1
	if product.ProductID > 0 {
		oldProduct, err := getProduct(product.ProductID)
		if err != nil {
			return addOrUpdateID, err
		}
		// if it exists, replace it. otherwise return error
		if oldProduct == nil {
			return 0, fmt.Errorf("product di [%d] does not exist", product.ProductID)
		}
		addOrUpdateID = product.ProductID
	} else {
		addOrUpdateID = getNextProductID()
		product.ProductID = addOrUpdateID
	}
	productMap.Lock()
	productMap.m[addOrUpdateID] = product
	productMap.Unlock()
	return addOrUpdateID, nil
} // addOrUpdateProduct

