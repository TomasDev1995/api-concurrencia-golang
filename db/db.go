package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// InitDB initializes the database connection and returns the *sql.DB object
func InitDB() (*sql.DB, error) {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load("./.env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	// Obtener las variables de entorno necesarias para la conexión a la base de datos
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Verificar que todas las variables de entorno están definidas
	if host == "" {
		return nil, errors.New("DB_HOST environment variable is not defined")
	}
	if port == "" {
		return nil, errors.New("DB_PORT environment variable is not defined")
	}
	if user == "" {
		return nil, errors.New("DB_USER environment variable is not defined")
	}
	if password == "" {
		return nil, errors.New("DB_PASSWORD environment variable is not defined")
	}
	if dbname == "" {
		return nil, errors.New("DB_NAME environment variable is not defined")
	}

	// Construir la cadena de conexión a la base de datos
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abrir la conexión a la base de datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Verificar si la base de datos está accesible
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}
