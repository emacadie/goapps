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
	// "github.com/jackc/pgx/v4"
	"strconv"
	"strings"
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
} // init

func getProduct(productID int) (*Product, error) {
	var funcName = fileNameData + "getProduct: "
	log.Printf( funcName + " with product ID %d \n", productID )

	ctx, cancel := context.WithTimeout(context.Background(), ( 15 * time.Second ))
	defer cancel()
	row := database.DbConn.QueryRow(
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
	_, err := database.DbConn.Query(ctx,  `delete from products where productid = $1`, productID )
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
	results, err := database.DbConn.Query(ctx,
		`select productId, manufacturer, sku, upc, to_char(pricePerUnit, '999D9'), 
        quantityOnHand, productName from products` )
	if err != nil {
		return nil, err
	}
	defer results.Close()
	products := make([]Product, 0)
	for results.Next() {
		var product Product
		fmt.Println( "Got a product" )
		err = results.Scan(&product.ProductID, 
			&product.Manufacturer, 
			&product.Sku, 
			&product.Upc, 
			&product.PricePerUnit, 
			&product.QuantityOnHand, 
			&product.ProductName)
		if err != nil {
			log.Print(funcName, "Got an error: ", err)
			log.Println( " " )
			log.Println( "Error was with product id " + strconv.Itoa( product.ProductID ) )
		}
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
		err := database.DbConn.QueryRow( ctx,  
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
	var funcName = fileNameData + "updateProduct: "
	log.Printf( "%s with product id %d \n", funcName, product.ProductID  )
	ctx, cancel := context.WithTimeout(context.Background(), ( 15 * time.Second ))
	defer cancel()
	_, err := database.DbConn.Exec( ctx,
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

func GetTopTenProducts() ( []Product, error ) {
	var funcName = fileNameData + "GetTopTenProducts: "
	ctx, cancel := context.WithTimeout( context.Background(), 3 * time.Second )
	defer cancel()
	results, err := database.DbConn.Query(ctx, 
		`select productId, manufacturer, sku, upc, pricePerUnit, quantityOnHand, productName
        from Products order by quantityOnHand desc limit 10`)
	if err != nil {
		log.Println( funcName + "error getting top 10 products from DB: " + err.Error() )
		return nil, err
	}
	defer results.Close()
	products := make( []Product, 0 )
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID, 
			&product.Manufacturer, 
			&product.Sku, 
			&product.Upc, 
			&product.PricePerUnit, 
			&product.QuantityOnHand, 
			&product.ProductName)
		products = append( products, product )
	} // for results.Next()
	return products, nil
} // GetTopTenProducts

func searchForProductData( productFilter ProductReportFilter ) ( []Product, error ) {
	var funcName = fileNameData + "searchForProductData: "
	ctx, cancel := context.WithTimeout( context.Background(), 3 * time.Second )
	defer cancel()
	
	var queryArgs = make([]interface{}, 0 )
	var queryBuilder strings.Builder
/*
	    queryBuilder.WriteString(`SELECT
        productId,
        LOWER(manufacturer),
        LOWER(sku),
        upc,
        pricePerUnit,
        quantityOnHand,
        LOWER(productName)    
        FROM products WHERE `)
*/

	queryBuilder.WriteString(
		`select productId, LOWER( manufacturer ), LOWER( sku ), upc, pricePerUnit, quantityOnHand, LOWER( productName ) 
        FROM products where `)

	if productFilter.NameFilter != "" {
		queryBuilder.WriteString( `productName like ? ` )
		queryArgs = append( queryArgs, "%" + strings.ToLower( productFilter.NameFilter) + "%" )
	}
	if productFilter.ManufacturerFilter != "" {
		if len( queryArgs ) > 0 {
			queryBuilder.WriteString( " and " )
		}
		queryBuilder.WriteString( `manufacturer like ?` )
		queryArgs = append( queryArgs, "%" + strings.ToLower( productFilter.ManufacturerFilter) + "%" )
	}
	if productFilter.SKUFilter != "" {
		if len( queryArgs ) > 0 {
			queryBuilder.WriteString( " and " )
		}
		queryBuilder.WriteString( `sku like ?` )
		queryArgs = append( queryArgs, "%" + strings.ToLower( productFilter.SKUFilter) + "%" )
	}
	results, err := database.DbConn.Query( ctx, queryBuilder.String(), queryArgs... )
	if err != nil {
		log.Println( funcName + "Error in query: " + err.Error() )
		return nil, err
	}
	defer results.Close()
	products := make( []Product, 0 )
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID, 
			&product.Manufacturer, 
			&product.Sku, 
			&product.Upc, 
			&product.PricePerUnit, 
			&product.QuantityOnHand, 
			&product.ProductName)
		products = append( products, product )
	} // for results.Next()
	return products, nil
} // searchForProductData

