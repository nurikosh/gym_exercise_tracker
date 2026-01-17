package main

import (
	"log"

	"gym/internal/storage/server"
)

func main() {
	db, err := server.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
