package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gomarket/routes"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env.example file: %v", err)
	}

	routes.Load()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
