package modelos

type Usuario struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Nombre     string `gorm:"size:50;not null" json:"nombre"`
	Contraseña string `gorm:"size:255;not null" json:"contraseña"` // Se debe hashear en el backend
	Id_rol     uint   `gorm:"not null" json:"id_rol" `
	Rol        Rol    `gorm:"foreignKey:id_rol"`
}
