package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
)

type service struct {
	db *sql.DB
}

var (
	dbname     = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service
)

// func ConnectDB() *service {

// 	if dbInstance != nil {
// 		return dbInstance
// 	}

// 	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	dbInstance = &service{
// 		db: db,
// 	}

// 	return dbInstance
// }

var DB *sqlx.DB

func ConnectDB() {
	var err error
	DB, err = sqlx.Connect("mysql", "username:password@tcp(127.0.0.1:3306)/blogdb")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dbname)
	return s.db.Close()
}
