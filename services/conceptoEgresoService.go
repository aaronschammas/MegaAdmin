package services

import (
	"MegaModa/modelos"

	"gorm.io/gorm"
)

// ConceptoEgresoService proporciona métodos para interactuar con el modelo ConceptoEgreso.
type ConceptoEgresoService struct {
	DB *gorm.DB
}

// NewConceptoEgresoService crea una nueva instancia de ConceptoEgresoService.
func NewConceptoEgresoService(db *gorm.DB) *ConceptoEgresoService {
	return &ConceptoEgresoService{DB: db}
}

// CrearConceptoEgreso agrega un nuevo concepto de egreso a la base de datos.
func (s *ConceptoEgresoService) CrearConceptoEgreso(concepto *modelos.ConceptoEgreso) error {
	return s.DB.Create(concepto).Error
}

// ObtenerConceptoEgreso obtiene un concepto de egreso por su ID.
func (s *ConceptoEgresoService) ObtenerConceptoEgreso(id int) (*modelos.ConceptoEgreso, error) {
	var concepto modelos.ConceptoEgreso
	err := s.DB.First(&concepto, id).Error
	return &concepto, err
}

// ActualizarConceptoEgreso actualiza un concepto de egreso existente.
func (s *ConceptoEgresoService) ActualizarConceptoEgreso(concepto *modelos.ConceptoEgreso) error {
	return s.DB.Save(concepto).Error
}

// EliminarConceptoEgreso elimina un concepto de egreso por su ID.
func (s *ConceptoEgresoService) EliminarConceptoEgreso(id int) error {
	return s.DB.Delete(&modelos.ConceptoEgreso{}, id).Error
}
