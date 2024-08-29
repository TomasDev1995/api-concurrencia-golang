package routes

import (
	"golang_with_docker_model_1/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/resource", controllers.GetResource).Methods("GET")
	router.HandleFunc("/api/v1/resource", controllers.CreateResource).Methods("POST")

	return router
}
