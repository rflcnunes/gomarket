package routes

import (
	"gomarket/controllers"
	"net/http"
)

func Load() {
	http.HandleFunc("/", controllers.GetAllProducts)
}
