package product

import (
	"fmt"
	"log"
	"time"
	"golang.org/x/net/websocket"
)

// const fileNameData = "In product.websocket."

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func productSocket( ws *websocket.Conn ) {
	var funcName = fileNameData + "productSocket: "
	done := make(chan struct{})
	log.Println( funcName + "new websocket connetion established" )
	go func( c *websocket.Conn ) {
		for {
			var msg message
			if err := websocket.JSON.Receive( ws, &msg ); err != nil {
				log.Printf( "%s error in websocket.JSON.Receieve: %s \n ", funcName, err )
				break
			}
			fmt.Printf( "received message %s \n", msg.Data )
			log.Printf( "%s received message %s \n", funcName, msg.Data )
			
		} // for
		close( done )
	} ( ws )
	loop:
	for {
		select {
		case <-done:
			log.Println( funcName + "Connection was closed, lets break out of here" )
			fmt.Println( "Connection was closed, lets break out of here" )
			break loop
		default:
			products, err := GetTopTenProducts()
			log.Println( funcName + "just called GetTopTenProducts()" )
			if err != nil {
				log.Printf( "%s error calling GetTopTenProducts: %s\n", funcName, err )
				break
			}
			if err := websocket.JSON.Send( ws, products ); err != nil {
				log.Printf( "%s error calling websocket.JSON.Send: %s \n ", funcName, err )
			}
			time.Sleep( 10 * time.Second )

		} // select

	} // for
	log.Println( funcName + "closing web socket" )
	defer ws.Close()
} // productSocket

