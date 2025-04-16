package services

import (
	"MegaModa/modelos"
	"fmt"

	"gorm.io/gorm"
)

// LocalService proporciona métodos para interactuar con el modelo Local.
type LocalService struct {
	DB *gorm.DB
}

// NewLocalService crea una nueva instancia de LocalService.
func NewLocalService(db *gorm.DB) *LocalService {
	return &LocalService{DB: db}
}

// CrearLocal agrega un nuevo local a la base de datos.
func (s *LocalService) CrearLocal(local *modelos.Local) error {
	return s.DB.Create(local).Error
}

// ObtenerLocal obtiene un local por su ID.
func (s *LocalService) ObtenerLocal(id int) (*modelos.Local, error) {
	var local modelos.Local
	err := s.DB.First(&local, id).Error
	if err != nil {
		return nil, err
	}
	return &local, nil
}

func (s *LocalService) ObtenerLocales() ([]modelos.Local, error) {
	var locales []modelos.Local
	err := s.DB.Find(&locales).Error
	return locales, err
}

// ActualizarLocal actualiza un local existente.
func (s *LocalService) ActualizarLocal(local *modelos.Local) error {
	return s.DB.Save(local).Error
}

// EliminarLocal elimina un local por su ID.
func (s *LocalService) EliminarLocal(id int) error {
	return s.DB.Delete(&modelos.Local{}, id).Error
}

func (s *LocalService) ObtenerLocalesConTurnos(fecha, turno string) ([]modelos.Local, error) {
	var locales []modelos.Local

	// Obtener los locales que tienen turnos en la fecha y turno especificados
	err := s.DB.Joins("JOIN turnos ON turnos.id_local = locals.id").
		Where("turnos.fecha = ?", fecha).
		Where("turnos.turno = ?", turno).
		Find(&locales).Error

	if err != nil {
		return nil, fmt.Errorf("error al obtener los locales con turnos: %v", err)
	}

	return locales, nil
}
