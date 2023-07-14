package main

import (
	"database/sql"
	"log"

	"github.com/after23/sharing-vision-be/api"
	db "github.com/after23/sharing-vision-be/db/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver      = "mysql"
	dbSource      = "root:secret@tcp(localhost:1357)/article"
	serverAddress = "0.0.0.0:1234"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Panicf("failed to connect to the database: %v", err)
	}
	q := db.New(conn)
	server := api.NewServer(q)
	
	err = server.Start(serverAddress)
	if err != nil {
		log.Panicf("failed to start the server: %v", err)
	}
}