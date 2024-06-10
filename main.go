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
var db *sql.DB

type Product struct {
	Id, Quantity      int
	Name, Description string
	Price             float64
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

func getAllProducts() []Product {
	rows, err := db.Query("SELECT id, name, description, quantity, price FROM products")

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	var products []Product

	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &quantity, &price)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Quantity = quantity
		product.Price = price

		products = append(products, product)
	}

	return products
}

func index(w http.ResponseWriter, r *http.Request) {
	products := getAllProducts()

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

	db = connectDB()
	defer db.Close()

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
