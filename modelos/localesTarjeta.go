package modelos

/*type LocalTarjeta struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	IDLocal   uint    `gorm:"not null"`
	Local     Local   `gorm:"foreignKey:IDLocal"`
	IDTarjeta uint    `gorm:"not null"`
	Tarjeta   Tarjeta `gorm:"foreignKey:IDTarjeta"`
}*/
type LocalTarjeta struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	IDLocal   uint    `gorm:"not null" json:"id_local"`
	IDTarjeta uint    `gorm:"not null" json:"id_tarjeta"`
	Local     Local   `gorm:"foreignKey:IDLocal" json:"-"`
	Tarjeta   Tarjeta `gorm:"foreignKey:IDTarjeta" json:"-"`
}
