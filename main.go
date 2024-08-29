package main

import (
	"log"
	"net/http"
	"os"

	"golang_with_docker_model_1/db"
	"golang_with_docker_model_1/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error cargando el archivo .env 'godotenv.Load()': %v", err)
	}

	// Inicializar la base de datos
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error inicializando la base de datos: %v", err)
	}
	defer dbConn.Close()

	// Inicializar las rutas
	router := routes.InitRoutes()

	// Configurar el servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Starting server on :%s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :%s: %v\n", port, err)
	}
}
