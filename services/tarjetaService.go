package services

import (
	"MegaModa/modelos"

	"gorm.io/gorm"
)

// TarjetaService proporciona métodos para interactuar con el modelo Tarjeta.
type TarjetaService struct {
	DB *gorm.DB
}

// NewTarjetaService crea una nueva instancia de TarjetaService.
func NewTarjetaService(db *gorm.DB) *TarjetaService {
	return &TarjetaService{DB: db}
}

// CrearTarjeta agrega una nueva tarjeta a la base de datos.
func (s *TarjetaService) CrearTarjeta(tarjeta *modelos.Tarjeta) error {
	return s.DB.Create(tarjeta).Error
}

// ObtenerTarjeta obtiene una tarjeta por su ID.
func (s *TarjetaService) ObtenerTarjeta(id int) (*modelos.Tarjeta, error) {
	var tarjeta modelos.Tarjeta
	err := s.DB.First(&tarjeta, id).Error
	return &tarjeta, err
}

// ActualizarTarjeta actualiza una tarjeta existente.
func (s *TarjetaService) ActualizarTarjeta(tarjeta *modelos.Tarjeta) error {
	return s.DB.Save(tarjeta).Error
}

// EliminarTarjeta elimina una tarjeta por su ID.
func (s *TarjetaService) EliminarTarjeta(id int) error {
	return s.DB.Delete(&modelos.Tarjeta{}, id).Error
}
