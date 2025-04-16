package DB

import (
	"gorm.io/gorm"
)

const (
	DBName = "MegaModa"
	DBUser = "root"
	DBPass = ""
	DBHost = "localhost"
	DBPort = "3306"
)

type Connection struct {
	DB *gorm.DB
}

const (
	TableSucursal       = "sucursales"
	TableTurno          = "turnos"
	TableVentaTarjeta   = "ventas_tarjetas"
	TableBilleteRendido = "billetes_rendidos"
)
