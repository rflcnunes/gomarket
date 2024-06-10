package main

import (
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
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
		log.Fatalf("Error loading .env file: %v", err)
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil).Error()
}
