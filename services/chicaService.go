package services

import (
	"MegaModa/modelos"

	"gorm.io/gorm"
)

// ChicaService proporciona métodos para interactuar con el modelo Chica.
type ChicaService struct {
	DB *gorm.DB
}

// NewChicaService crea una nueva instancia de ChicaService.
func NewChicaService(db *gorm.DB) *ChicaService {
	return &ChicaService{DB: db}
}

// CrearChica agrega una nueva chica a la base de datos.
func (s *ChicaService) CrearChica(chica *modelos.Chica) error {
	return s.DB.Create(chica).Error
}

// ObtenerChica obtiene una chica por su ID.
func (s *ChicaService) ObtenerChica(id int) (*modelos.Chica, error) {
	var chica modelos.Chica
	err := s.DB.First(&chica, id).Error
	return &chica, err
}

// ActualizarChica actualiza una chica existente.
func (s *ChicaService) ActualizarChica(chica *modelos.Chica) error {
	return s.DB.Save(chica).Error
}

// EliminarChica elimina una chica por su ID.
func (s *ChicaService) EliminarChica(id int) error {
	return s.DB.Delete(&modelos.Chica{}, id).Error
}
