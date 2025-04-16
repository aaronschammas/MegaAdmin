package modelos

// Tarjeta representa la tabla de tarjetas.
type Tarjeta struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Tipo   string `gorm:"type:enum('Débito', 'Crédito');not null"`
	Nombre string `gorm:"size:50;not null"`
}
