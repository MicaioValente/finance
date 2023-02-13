package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/micaiovalente/finance/api"
	db "github.com/micaiovalente/finance/db/sqlc"
	"log"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	dbDriver := "postgres"
	dbSource := "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress := "0.0.0.0:8000"
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api: ", err)
	}
}
