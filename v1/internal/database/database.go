package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

// var DB *sqlx.DB
var DB *sql.DB

func ConnectDB() *sql.DB {

	var err error
	DB, err = sql.Open("mysql", os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
	return DB

}
