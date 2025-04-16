package services

import (
	"MegaModa/modelos"

	"gorm.io/gorm"
)

// PermisoService proporciona métodos para interactuar con el modelo Permiso.
type PermisoService struct {
	DB *gorm.DB
}

// NewPermisoService crea una nueva instancia de PermisoService.
func NewPermisoService(db *gorm.DB) *PermisoService {
	return &PermisoService{DB: db}
}

// CrearPermiso agrega un nuevo permiso a la base de datos.
func (s *PermisoService) CrearPermiso(permiso *modelos.Permiso) error {
	return s.DB.Create(permiso).Error
}

// ObtenerPermiso obtiene un permiso por su ID.
func (s *PermisoService) ObtenerPermiso(id int) (*modelos.Permiso, error) {
	var permiso modelos.Permiso
	err := s.DB.First(&permiso, id).Error
	return &permiso, err
}

// ActualizarPermiso actualiza un permiso existente.
func (s *PermisoService) ActualizarPermiso(permiso *modelos.Permiso) error {
	return s.DB.Save(permiso).Error
}

// EliminarPermiso elimina un permiso por su ID.
func (s *PermisoService) EliminarPermiso(id int) error {
	return s.DB.Delete(&modelos.Permiso{}, id).Error
}

// ObtenerPermisosPorIDUsuario obtiene todos los permisos de un usuario
func (s *PermisoService) ObtenerPermisosPorIDUsuario(idUsuario uint) ([]modelos.Permiso, error) {
	var permisos []modelos.Permiso
	err := s.DB.Where("id_usuario = ?", idUsuario).Find(&permisos).Error
	return permisos, err
}
