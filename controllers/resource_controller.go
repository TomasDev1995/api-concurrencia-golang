package controllers

import (
	"encoding/json"
	"net/http"
)

func GetResource(w http.ResponseWriter, r *http.Request) {
	// Ejemplo de respuesta
	response := map[string]string{"message": "Resource fetched successfully"}
	json.NewEncoder(w).Encode(response)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para crear el recurso
	response := map[string]string{"message": "Resource created successfully"}
	json.NewEncoder(w).Encode(response)
}
