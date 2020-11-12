package product

import (
	"context"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	// "os"
	"sort"
	"database/sql"
	"strconv"
	"sync"
	"time"

	"shelfunit.info/golang/inventoryservice/database"
)

const fileNameData = "In product.data."

var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{m: make(map[int]Product)}

func init() {
	var funcName = fileNameData + "init: "
	log.Println( funcName + "starting product.data.go" )
	/*
	fmt.Println("loading products..")
	prodMap, err := loadProductMap()
	productMap.m = prodMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d products loaded \n", len(productMap.m))
	*/
} // init
/*
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
*/

func getProduct(productID int) (*Product, error) {
	var funcName = fileNameData + "getProduct: "
	log.Printf( funcName + " with product ID %d \n", productID )

	ctx, cancel := context.WithTimeout(context.Background(), ( 15 * time.Second ))
	defer cancel()
	row := database.DbConn.QueryRowContext(
		ctx,
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

} // getProduct

func removeProduct(productID int) error {
	var funcName = fileNameData + "removeProduct: "
	log.Printf( "In %s with productID: %d \n", funcName, productID )
	ctx, cancel := context.WithTimeout(context.Background(), ( 15 * time.Second ))
	defer cancel()
	_, err := database.DbConn.QueryContext(ctx,  `delete from products where productid = $1`, productID )
	if err != nil {
		log.Print( funcName, "Got an err: ", err )
		log.Println( " " )
		return err
	}
	return nil
}

func getProductList() ([]Product, error) {
	var funcName = fileNameData + "getProductList: "
	log.Println( funcName )
	ctx, cancel := context.WithTimeout(context.Background(), ( 15 * time.Second ))
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx,
		`select productId, manufacturer, sku, upc, pricePerUnit, quantityOnHand, productName from products` )
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

func insertProduct( product Product ) ( int, error ) {
	var funcName = fileNameData + "insertProduct: "
	log.Println( funcName )

	ctx, cancel := context.WithTimeout(context.Background(), ( 15 * time.Second ))
	defer cancel()
	// using QueryRow instead of Exec since the driver I was using does not implement LastInsertId
	newProductID := 0
		err := database.DbConn.QueryRowContext( ctx,  
		`insert into products (manufacturer, sku, upc, pricePerUnit, quantityOnHand, productName)
        values ( $1, $2, $3, $4, $5, $6 ) returning productid`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName	).Scan(&newProductID)

	if err != nil {
		log.Print( funcName, "Got an err: ", err )
		log.Println( " " )
		return 0, err
	}

	log.Printf( "%s: Adding id %d \n", funcName, newProductID )
	return int( newProductID ), nil
} // insertProduct


func updateProduct( product Product ) error {
	var funcName = "In " + fileNameData + "updateProduct: "
	log.Printf( "%s with product id %d \n", funcName, product.ProductID  )
	ctx, cancel := context.WithTimeout(context.Background(), ( 15 * time.Second ))
	defer cancel()
	_, err := database.DbConn.ExecContext( ctx,
		`update products set manufacturer = $1,
         sku = $2,
         upc = $3,
         pricePerUnit = CAST ( $4 AS numeric( 13, 2 ) ),
         quantityOnHand = $5,
         productName = $6
         where productid = $7`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID)
	if err != nil {
		log.Print( funcName, "Got an err: ", err )
		log.Println( " " )
		return err
	}
	return nil
	
} // updateProduct


