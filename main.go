package main

import (
	"log"
	"net/http"

	"routes" // Importa tu paquete de rutas
)

func main() {
	// Configura el servidor HTTP
	router := routes.InitRoutes() // Inicia las rutas desde un paquete separado
	server := &http.Server{
		Addr:    ":8080", // Escucha en el puerto 8080
		Handler: router,  // Usa el router configurado
	}

	// Inicia el servidor y maneja errores potenciales
	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}
}
