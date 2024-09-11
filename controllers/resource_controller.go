package controllers

import (
	"encoding/json"
	"net/http"

	service "golang_with_docker_model_1/services" // Ajusta según la estructura de tu proyecto
)

// ResourceHandler maneja las solicitudes relacionadas con recursos
type ResourceHandler struct {
	service service.ResourceService
}

// NewResourceHandler crea una nueva instancia de ResourceHandler
func NewResourceHandler(service service.ResourceService) *ResourceHandler {
	return &ResourceHandler{service: service}
}

// CreateResource maneja la solicitud para crear un recurso
func (h *ResourceHandler) CreateResource(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Extract the data from the request
	name := requestData["name"]
	email := requestData["email"]

	// Call the service to create the resource
	if err := h.service.CreateResource(name, email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a success response
	response := map[string]string{"message": "Resource created successfully"}
	json.NewEncoder(w).Encode(response)
}
