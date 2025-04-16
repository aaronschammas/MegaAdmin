package modelos

type Local struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Nombre string `gorm:"size:50;not null"`
}

func (Local) TableName() string {
	return "locals"
}
