package modelos

type Permiso struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	IDUsuario uint    `gorm:"not null"`
	Usuario   Usuario `gorm:"foreignKey:IDUsuario"`
	IDLocal   uint    `gorm:"not null"`
	Local     Local   `gorm:"foreignKey:IDLocal"`
}
