package routes

import (
	"gomarket/controllers"
	"net/http"
)

func ProductsRoutesLoad() {
	http.HandleFunc("/products/create", controllers.CreateProduct)
	http.HandleFunc("/products/insert", controllers.GetAllProducts)
	http.HandleFunc("/products/delete", controllers.DeleteProduct)
}
