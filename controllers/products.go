package controllers

import (
	"gomarket/models"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	err := templates.ExecuteTemplate(w, "Index", products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
