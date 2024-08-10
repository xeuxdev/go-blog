package db

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

func main() {
	m, err := migrate.New(
		"./create-tables.sql",
		os.Getenv("DB_URL"))
	m.Steps(2)
}
