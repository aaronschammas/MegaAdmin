package modelos

// ConceptoEgreso representa la tabla de conceptos de egresos.
type ConceptoEgreso struct {
	ID       uint    `gorm:"primaryKey;autoIncrement"`
	Concepto string  `gorm:"size:255;not null"`
	Monto    float64 `gorm:"not null"`
	IDTurno  uint    `gorm:"column:id_turno;not null" json:"id_turno"`
	Turno    Turno   `gorm:"foreignKey:IDTurno" json:"-"`
}

func (ConceptoEgreso) TableName() string {
	return "ConceptoEgresos" // Cambia según el nombre de tu tabla
}
