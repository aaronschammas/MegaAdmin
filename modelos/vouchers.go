package modelos

type Voucher struct {
	ID      uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Tipo    string  `gorm:"type:enum('Crédito', 'Débito');not null" json:"tipo"` // "Crédito" o "Débito"
	Tarjeta string  `gorm:"size:50;not null" json:"tarjeta"`                     // Nombre de la tarjeta (Visa, MasterCard, etc.)
	Cuotas  int     `gorm:"not null" json:"cuotas"`                              // Número de cuotas
	Monto   float64 `gorm:"not null" json:"monto"`                               // Monto asociado al voucher
	IDTurno uint    `gorm:"column:id_turno;not null" json:"id_turno"`            // ID del turno relacionado
	Turno   Turno   `gorm:"foreignKey:IDTurno" json:"-"`                         // Relación con el turno
}
