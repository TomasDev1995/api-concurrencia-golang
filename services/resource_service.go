package service

import (
	repository "golang_with_docker_model_1/repositories" // Ajusta según la estructura de tu proyecto
)

// ResourceService es una interfaz que define los métodos de servicio para recursos
type ResourceService interface {
	CreateResource(name, email string) error
}

// AppResourceService implementa la interfaz ResourceService
type AppResourceService struct {
	repo repository.ResourceRepository
}

// NewAppResourceService crea una nueva instancia de AppResourceService
func NewAppResourceService(repo repository.ResourceRepository) *AppResourceService {
	return &AppResourceService{repo: repo}
}

// CreateResource delega la creación del recurso al repositorio
func (s *AppResourceService) CreateResource(name, email string) error {
	return s.repo.CreateResource(name, email)
}
