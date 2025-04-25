package main

import (
	"context"
	"ecommerce-go-api/api"
	db "ecommerce-go-api/db/sqlc"
	"ecommerce-go-api/util"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Could not set up new server and routing: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	fmt.Println(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

}
