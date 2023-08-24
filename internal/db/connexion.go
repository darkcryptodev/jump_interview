package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	// Open postgresql db
	db, err := sql.Open("postgres", "host=localhost port=5432 user=jump password=password dbname=jump sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Check connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully!")
	return db, nil
}
