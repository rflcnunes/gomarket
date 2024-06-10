package controllers

import (
	"gomarket/models"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	err := templates.ExecuteTemplate(w, "Index", products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		quantity := r.FormValue("quantity")
		price := r.FormValue("price")

		convertedQuantity, err := strconv.Atoi(quantity)
		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		product := models.Product{
			Name:        name,
			Description: description,
			Quantity:    convertedQuantity,
			Price:       convertedPrice,
		}

		models.CreateProduct(product)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := templates.ExecuteTemplate(w, "Products/Create", nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
