package database

import (
	// "database/sql"
	"log"
	// "time"

	"github.com/jackc/pgx/v4"
	"context"

)

// var DbConn *sql.DB // orig
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


	// orig
	// DbConn, err = sql.Open( "postgres", "postgres://go_web_usr:this-is-twitter@localhost:5433/inventorydb?sslmode=disable" )
	/*
	"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	psql --port=5433 -h localhost -U go_web_usr -d inventorydb
    */
	if err != nil {
		log.Fatal( err )
	}
	// DbConn.SetMaxOpenConns( 10 )
	// DbConn.SetMaxIdleConns( 10 )
	// DbConn.SetConnMaxLifetime( 60 * time.Second )
}

