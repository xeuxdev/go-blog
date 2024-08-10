package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func ConnectDB() *sql.DB {
	dsn := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
	return db
}
