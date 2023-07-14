package main

import (
	"database/sql"
	"log"

	"github.com/after23/sharing-vision-be/api"
	db "github.com/after23/sharing-vision-be/db/sqlc"
	"github.com/after23/sharing-vision-be/util"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	util.LoadEnv(".")
}

func main() {
	conn, err := sql.Open(util.GetConfig().DBDriver, util.GetConfig().DBSource)
	if err != nil {
		log.Panicf("failed to connect to the database: %v", err)
	}
	defer conn.Close()
	q := db.New(conn)
	server := api.NewServer(q)
	
	err = server.Start(util.GetConfig().ServerAddress)
	if err != nil {
		log.Panicf("failed to start the server: %v", err)
	}
}