package modelos

type Chica struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Nombre  string `gorm:"size:50;not null"`
	IDTurno uint   `gorm:"not null"`
	Turno   Turno  `gorm:"foreignKey:IDTurno"`
}
