package routes

import (
	"golang_with_docker_model_1/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes(handler *controllers.ResourceHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/resource", handler.CreateResource).Methods("POST")

	return router
}
