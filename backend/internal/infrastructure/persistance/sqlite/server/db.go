package server

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

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

func RunMigrations(db *sql.DB) error {
	path := filepath.Join("internal/storage/migrations", "001_init.sql")

	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error while reading: %v", err)
	}

	_, err = db.Exec(string(file))
	if err != nil {
		return fmt.Errorf("Migration error: %v", err)
	}

	return nil
}
