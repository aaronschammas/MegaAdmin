package services

import (
	"MegaModa/modelos"
	"fmt"

	"gorm.io/gorm"
)

type TurnoConChicas struct {
	modelos.Turno
	Chicas []modelos.Chica `json:"chicas"`
}

// TurnoService proporciona métodos para interactuar con el modelo Turno.
type TurnoService struct {
	DB *gorm.DB
}

// NewTurnoService crea una nueva instancia de TurnoService.
func NewTurnoService(db *gorm.DB) *TurnoService {
	return &TurnoService{DB: db}
}

// CrearTurno agrega un nuevo turno a la base de datos.
func (s *TurnoService) CrearTurno(turno *modelos.Turno) error {
	return s.DB.Create(turno).Error
}

// ObtenerTurno obtiene un turno por su ID.
func (s *TurnoService) ObtenerTurno(id int) (*modelos.Turno, error) {
	var turno modelos.Turno
	err := s.DB.First(&turno, id).Error
	return &turno, err
}

// Sacar todos los turnos
func (s *TurnoService) ObtenerTurnos() ([]modelos.Turno, error) {
	var turnos []modelos.Turno
	err := s.DB.Find(&turnos).Error
	return turnos, err
}

// ActualizarTurno actualiza un turno existente.
func (s *TurnoService) ActualizarTurno(turno *modelos.Turno) error {
	return s.DB.Save(turno).Error
}

// EliminarTurno elimina un turno por su ID.
func (s *TurnoService) EliminarTurno(id int) error {
	return s.DB.Delete(&modelos.Turno{}, id).Error
}

func (s *TurnoService) ObtenerTurnosNull() ([]modelos.Turno, error) {
	var turnos []modelos.Turno
	err := s.DB.Where("turno IS NULL").Find(&turnos).Error
	return turnos, err
}

func (s *TurnoService) FiltrarTurnos(fecha, turno, localID string) ([]modelos.Turno, error) {
	var turnos []modelos.Turno
	query := s.DB.Where("fecha = ?", fecha).Where("turno = ?", turno)

	if localID != "" {
		query = query.Where("id_local = ?", localID)
	}

	err := query.Find(&turnos).Error
	return turnos, err
}

// Sacar los locales con la relacion existente entre la tabla turnos y el los locales

func (s *TurnoService) ObtenerTurnoPorIDLocal(idLocal uint) (*modelos.Turno, error) {
	var turno modelos.Turno
	err := s.DB.Joins("JOIN locals ON turnos.id_local = locals.id").
		Where("locals.id = ?", idLocal).
		Select("turnos.id, turnos.fecha, turnos.turno, turnos.efectivo, turnos.tarjeta_credito, turnos.tarjeta_debito, turnos.financiera, turnos.total_egresos, turnos.total_ventas, turnos.diferencia_caja, turnos.id_local, turnos.id_usuario, turnos.timestamp").
		First(&turno).Error
	return &turno, err
}

// Obtener turnos junto con las chicas asociadas para un local específico
func (s *TurnoService) ObtenerTurnosConChicas(idLocal uint) ([]TurnoConChicas, error) {
	var turnos []modelos.Turno
	var turnosConChicas []TurnoConChicas

	// Obtener los turnos para el local específico
	err := s.DB.Where("id_local = ?", idLocal).Find(&turnos).Error
	if err != nil {
		return nil, fmt.Errorf("error al obtener los turnos: %v", err)
	}

	// Para cada turno, obtener las chicas asociadas
	for _, turno := range turnos {
		var chicas []modelos.Chica
		err := s.DB.Where("id_turno = ?", turno.ID).Find(&chicas).Error
		if err != nil {
			return nil, fmt.Errorf("error al obtener las chicas para el turno %d: %v", turno.ID, err)
		}

		// Crear el struct temporal con el turno y las chicas
		turnoConChicas := TurnoConChicas{
			Turno:  turno,
			Chicas: chicas,
		}

		turnosConChicas = append(turnosConChicas, turnoConChicas)
	}

	return turnosConChicas, nil
}

// otro service de turno flojo de papeles
func (s *TurnoService) ObtenerTurnosConChicasYFechas(idLocal uint, fechaDesde, fechaHasta, turno string) ([]TurnoConChicas, error) {
	var turnos []modelos.Turno
	var turnosConChicas []TurnoConChicas

	// Construir la consulta base
	query := s.DB.Where("fecha BETWEEN ? AND ?", fechaDesde, fechaHasta)

	// Filtrar por local si se proporciona un ID específico
	if idLocal != 0 {
		query = query.Where("id_local = ?", idLocal)
	}

	// Filtrar por turno si se proporciona (Mañana/Tarde)
	if turno != "" {
		query = query.Where("turno = ?", turno)
	}

	// Ejecutar la consulta para obtener los turnos
	err := query.Find(&turnos).Error
	if err != nil {
		return nil, fmt.Errorf("error al obtener los turnos: %v", err)
	}

	// Para cada turno, obtener las chicas asociadas
	for _, turno := range turnos {
		var chicas []modelos.Chica
		err := s.DB.Where("id_turno = ?", turno.ID).Find(&chicas).Error
		if err != nil {
			return nil, fmt.Errorf("error al obtener las chicas para el turno %d: %v", turno.ID, err)
		}

		// Crear el struct temporal con el turno y las chicas
		turnoConChicas := TurnoConChicas{
			Turno:  turno,
			Chicas: chicas,
		}

		// Agregar el struct a la lista final
		turnosConChicas = append(turnosConChicas, turnoConChicas)
	}

	return turnosConChicas, nil
}

// ta medio rara el modelo de esta funcion si me ocurre algo mejor lo hago pero es lo mismo del otro pero funciona llegandole un array de ese modelo
type TurnosModificados struct {
	ID    uint   `json:"id"`
	Turno string `json:"turno"`
}

func (s *TurnoService) ActualizarTurnos(turnos []TurnosModificados) error {
	for _, t := range turnos {
		err := s.DB.Model(&modelos.Turno{}).
			Where("id = ?", t.ID).
			Update("turno", t.Turno).Error

		if err != nil {
			return fmt.Errorf("error al actualizar el turno con ID %d: %v", t.ID, err)
		}
	}
	return nil
}
