package main

import (
	"database/sql"
	"fmt"
	"healthyapp/api"
	db "healthyapp/db/sqlc"
	"healthyapp/utils"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func main() {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error Config"))
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error connect database"))
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed start server"))
	}
	server.Start(config.HTTPServerAddress)
}
