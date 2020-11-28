package database

import (
	"log"
	"github.com/jackc/pgx/v4"
	"context"

)

var DbConn *pgx.Conn



func SetupDatabase() {
	var err error
	config, err := pgx.ParseConfig( "postgres://go_web_usr:this-is-twitter@localhost:5433/inventorydb?sslmode=disable"  )
    if err != nil {
        log.Fatal("error configuring the database: ", err)
    }

	DbConn, err = pgx.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

	if err != nil {
		log.Fatal( err )
	}
}

