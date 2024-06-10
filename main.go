package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gomarket/models"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

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

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
