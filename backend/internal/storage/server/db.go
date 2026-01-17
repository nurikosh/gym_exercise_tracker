package server

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName     = "gym.db"
	driverName = "sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open(driverName, dbName)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database %v", err)
	}

	return db, nil
}

func Run