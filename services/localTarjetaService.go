package services

import (
	"MegaModa/modelos"
	"fmt"

	"gorm.io/gorm"
)

// LocalTarjetaService proporciona métodos para interactuar con el modelo LocalTarjeta.
type LocalTarjetaService struct {
	DB *gorm.DB
}

// NewLocalTarjetaService crea una nueva instancia de LocalTarjetaService.
func NewLocalTarjetaService(db *gorm.DB) *LocalTarjetaService {
	return &LocalTarjetaService{DB: db}
}

// CrearLocalTarjeta agrega una nueva relación entre local y tarjeta a la base de datos.
func (s *LocalTarjetaService) CrearLocalTarjeta(localTarjeta *modelos.LocalTarjeta) error {
	return s.DB.Create(localTarjeta).Error
}

// ObtenerLocalTarjeta obtiene una relación entre local y tarjeta por su ID.
func (s *LocalTarjetaService) ObtenerLocalTarjeta(id int) (*modelos.LocalTarjeta, error) {
	var localTarjeta modelos.LocalTarjeta
	err := s.DB.First(&localTarjeta, id).Error
	return &localTarjeta, err
}

// ActualizarLocalTarjeta actualiza una relación entre local y tarjeta existente.
func (s *LocalTarjetaService) ActualizarLocalTarjeta(localTarjeta *modelos.LocalTarjeta) error {
	return s.DB.Save(localTarjeta).Error
}

// EliminarLocalTarjeta elimina una relación entre local y tarjeta por su ID.
func (s *LocalTarjetaService) EliminarLocalTarjeta(id int) error {
	return s.DB.Delete(&modelos.LocalTarjeta{}, id).Error
}

// Sacar las tarjetas que tiene un local segun el ID
func (s *LocalTarjetaService) ObtenerTarjetasPorLocal(idLocal uint) ([]modelos.Tarjeta, error) {
	var tarjetas []modelos.Tarjeta

	// Realizar una consulta JOIN para obtener los detalles completos de las tarjetas
	err := s.DB.Joins("JOIN local_tarjeta ON local_tarjeta.id_tarjeta = tarjeta.id").
		Where("local_tarjeta.id_local = ?", idLocal).
		Select("tarjeta.id, tarjeta.tipo, tarjeta.nombre"). // Seleccionar los campos necesarios
		Find(&tarjetas).Error

	if err != nil {
		return nil, fmt.Errorf("error al obtener las tarjetas del local: %v", err)
	}

	return tarjetas, nil
}
