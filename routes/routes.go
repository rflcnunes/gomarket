package routes

import (
	"gomarket/controllers"
	"net/http"
)

func Load() {
	http.HandleFunc("/", controllers.GetAllProducts)
	http.HandleFunc("/products/create", controllers.CreateProduct)
	http.HandleFunc("/products/insert", controllers.GetAllProducts)
	http.HandleFunc("/products/delete", controllers.DeleteProduct)
}
