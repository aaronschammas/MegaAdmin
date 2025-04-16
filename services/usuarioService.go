package services

import (
	"MegaModa/modelos"

	"gorm.io/gorm"
)

// UsuarioService proporciona métodos para interactuar con el modelo Usuario.
type UsuarioService struct {
	DB *gorm.DB
}

// NewUsuarioService crea una nueva instancia de UsuarioService.
func NewUsuarioService(db *gorm.DB) *UsuarioService {
	return &UsuarioService{DB: db}
}

// CrearUsuario agrega un nuevo usuario a la base de datos.
func (s *UsuarioService) CrearUsuario(usuario *modelos.Usuario) error {
	return s.DB.Create(usuario).Error
}

// ObtenerUsuario obtiene un usuario por su ID.
func (s *UsuarioService) ObtenerUsuario(id int) (*modelos.Usuario, error) {
	var usuario modelos.Usuario
	err := s.DB.First(&usuario, id).Error
	return &usuario, err
}
func (s *UsuarioService) ObtenerUsuarioPorNombre(nombre string) (modelos.Usuario, error) {
	var usuario modelos.Usuario
	err := s.DB.Where("nombre = ?", nombre).First(&usuario).Error
	return usuario, err
}

// ActualizarUsuario actualiza un usuario existente.
func (s *UsuarioService) ActualizarUsuario(usuario *modelos.Usuario) error {
	return s.DB.Save(usuario).Error
}

// EliminarUsuario elimina un usuario por su ID.
func (s *UsuarioService) EliminarUsuario(id int) error {
	return s.DB.Delete(&modelos.Usuario{}, id).Error
}
