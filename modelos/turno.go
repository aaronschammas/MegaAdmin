package modelos

import (
	"time"

	"gorm.io/gorm"
)

/*
	type Turno struct {
		ID             uint      `gorm:"primaryKey;autoIncrement"`
		Fecha          time.Time `gorm:"not null;default:CURRENT_DATE"`
		Turno          string    `gorm:"type:enum('Mañana', 'Tarde');default:null"`
		Efectivo       float64   `gorm:"not null"`
		TarjetaCredito float64   `gorm:"not null"`
		TarjetaDebito  float64   `gorm:"not null"`
		Financiera     float64   `gorm:"not null"`
		TotalEgresos   float64   `gorm:"not null"`
		TotalVentas    float64   `gorm:"<-:false;column:total_ventas"` // Calculado automáticamente
		DiferenciaCaja float64   `gorm:"not null"`
		IDLocal        uint      `gorm:"not null"`
		Local          Local     `gorm:"foreignKey:IDLocal"`
		IDUsuario      uint      `gorm:"not null"`
		Usuario        Usuario   `gorm:"foreignKey:IDUsuario"`
		Timestamp      time.Time `gorm:"autoCreateTime"`
	}
*/
type Turno struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Fecha          time.Time `gorm:"not null;default:CURRENT_DATE" json:"fecha"`
	Turno          string    `gorm:"type:enum('Mañana', 'Tarde');default:null" json:"tipo"`
	Efectivo       float64   `gorm:"not null" json:"ingreso_ef"`
	TarjetaCredito float64   `gorm:"not null" json:"ingreso_credito"`
	TarjetaDebito  float64   `gorm:"not null" json:"ingreso_debito"`
	Financiera     float64   `gorm:"not null" json:"financiera"`
	TotalEgresos   float64   `gorm:"not null" json:"egreso_total"`
	TotalVentas    float64   `gorm:"column:total_ventas" json:"total_ventas"` // Calculado automáticamente
	DiferenciaCaja float64   `gorm:"not null" json:"diferencia_caja"`
	IDLocal        uint      `gorm:"not null" json:"id_local"`
	Local          Local     `gorm:"foreignKey:IDLocal" json:"-"`
	IDUsuario      uint      `gorm:"not null" json:"id_usuario"`
	Usuario        Usuario   `gorm:"foreignKey:IDUsuario" json:"-"`
	Timestamp      time.Time `gorm:"autoCreateTime" json:"timestamp"`
}

// Método para calcular `TotalVentas` automáticamente
func (t *Turno) BeforeSave(tx *gorm.DB) (err error) {
	t.TotalVentas = t.Efectivo + t.TarjetaCredito + t.TarjetaDebito + t.Financiera
	return
}
