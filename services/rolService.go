package services

import (
	"MegaModa/modelos"

	"gorm.io/gorm"
)

// RolService proporciona métodos para interactuar con el modelo Rol.
type RolService struct {
	DB *gorm.DB
}

// NewRolService crea una nueva instancia de RolService.
func NewRolService(db *gorm.DB) *RolService {
	return &RolService{DB: db}
}

// CrearRol agrega un nuevo rol a la base de datos.
func (s *RolService) CrearRol(rol *modelos.Rol) error {
	return s.DB.Create(rol).Error
}

// ObtenerRol obtiene un rol por su ID.
func (s *RolService) ObtenerRol(id int) (*modelos.Rol, error) {
	var rol modelos.Rol
	err := s.DB.First(&rol, id).Error
	return &rol, err
}

// ActualizarRol actualiza un rol existente.
func (s *RolService) ActualizarRol(rol *modelos.Rol) error {
	return s.DB.Save(rol).Error
}

// EliminarRol elimina un rol por su ID.
func (s *RolService) EliminarRol(id int) error {
	return s.DB.Delete(&modelos.Rol{}, id).Error
}
