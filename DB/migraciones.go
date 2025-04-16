package DB

import (
	"MegaModa/modelos"
	"log"
)

func Migrate(conn *Connection) {
	err := conn.DB.AutoMigrate(
		&modelos.Rol{},            // Modelo de roles
		&modelos.Usuario{},        // Modelo de usuarios
		&modelos.Local{},          // Modelo de locales
		&modelos.Tarjeta{},        // Modelo de tarjetas
		&modelos.LocalTarjeta{},   // Relación muchos a muchos entre locales y tarjetas
		&modelos.Turno{},          // Modelo de turnos
		&modelos.Chica{},          // Modelo de chicas
		&modelos.ConceptoEgreso{}, // Modelo de conceptos de egresos
		&modelos.Permiso{},        // Modelo de permisos
		&modelos.Voucher{},        // Modelo de vouchers
	)
	if err != nil {
		log.Fatalf("Error durante la migración: %v", err)
	}
	log.Println("Migración completada con éxito")
}
