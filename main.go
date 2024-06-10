package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

func connectDB() *sql.DB {
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

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Product 1", "Description 1", 1.99, 10},
		{"Product 2", "Description 2", 2.99, 20},
		{"Product 3", "Description 3", 3.99, 30},
		{"Product 4", "Description 4", 4.99, 40},
	}

	err := templates.ExecuteTemplate(w, "Index", products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env.example file: %v", err)
	}

	db := connectDB()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil).Error()
}
