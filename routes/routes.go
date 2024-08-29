package routes

import (
	"api-concurrencia-golang/controllers"
	"gorilla/mux" // Ejemplo usando Gorilla Mux como router
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define tus rutas aquí
	router.HandleFunc("/api/v1/resource", controllers.GetResource).Methods("GET")
	router.HandleFunc("/api/v1/resource", controllers.CreateResource).Methods("POST")

	return router
}
