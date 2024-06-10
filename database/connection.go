package database

import (
	"database/sql"
	"os"
)

func ConnectDB() *sql.DB {
	userName := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_DB_HOST")
	dbName := os.Getenv("POSTGRES_DB")

	connectionString := "user=" + userName + " password=" + password + " host=" + host + " dbname=" + dbName + " sslmode=disable"

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return db
}
