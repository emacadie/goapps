package database

import (
	"database/sql"
	"log"
)

var DbConn *sql.DB

func SetupDatabase() {
	var err error
	DbConn, err = sql.Open( "postgres", "postgres://go_web_usr:this-is-twitter@localhost:5433/inventorydb?sslmode=disable" )
	/*
	"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	"postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable"
	psql --port=5433 -h localhost -U go_web_usr -d inventorydb
    */
	if err != nil {
		log.Fatal( err )
	}
}

