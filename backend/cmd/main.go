package main

import (
	"fmt"
	"gym/internal/infrastructure/repository/sqlite/server"
	"log"
)

func main() {
	db, err := server.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = server.RunMigrations(db)
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	fmt.Print("Success")
}
