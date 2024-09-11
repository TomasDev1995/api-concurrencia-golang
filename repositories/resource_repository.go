package repository

import (
	"database/sql"
	"fmt"
)

// ResourceRepository es una interfaz que define los métodos de acceso a datos para recursos
type ResourceRepository interface {
	CreateResource(name, email string) error
}

// PostgresResourceRepository implementa la interfaz ResourceRepository usando PostgreSQL
type PostgresResourceRepository struct {
	db *sql.DB
}

// NewPostgresResourceRepository crea una nueva instancia de PostgresResourceRepository
func NewPostgresResourceRepository(db *sql.DB) *PostgresResourceRepository {
	return &PostgresResourceRepository{db: db}
}

// CreateResource inserta un nuevo recurso en la base de datos
func (r *PostgresResourceRepository) CreateResource(name, email string) error {
	_, err := r.db.Exec("INSERT INTO prueba (nombre, email) VALUES ($1, $2)", name, email)
	if err != nil {
		return fmt.Errorf("error inserting resource into the database: %w", err)
	}
	return nil
}
