package DB

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB // Conexión global

func NewConnection() (*Connection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DBUser, DBPass, DBHost, DBPort, DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		if strings.Contains(err.Error(), "Unknown database") {
			db, err = crearBaseDeDatosYRealizarMigracion()
			if err != nil {
				return nil, fmt.Errorf("error conectando a MySQL: %v", err)
			}
		} else {
			return nil, fmt.Errorf("error conectando a MySQL: %v", err)
		}
	}

	// Asigna la conexión a la variable global
	GlobalDB = db

	return &Connection{DB: db}, nil
}

func crearBaseDeDatosYRealizarMigracion() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", DBUser, DBPass, DBHost, DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Crear la base de datos
	db.Exec("CREATE DATABASE " + DBName)
	// Realizar la migración
	Migrate(&Connection{DB: db})
	return db, nil
}
