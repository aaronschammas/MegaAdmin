package modelos

type Rol struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Nombre string `gorm:"size:50;unique;not null"`
}
