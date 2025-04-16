package main

import (
	"MegaModa/DB"   // Asegúrate de importar correctamente el paquete
	"MegaModa/mapa" // Donde configuras las rutas
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Establecer conexión a la base de datos
	connection, err := DB.NewConnection()
	if err != nil {

		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer func() {
		sqlDB, _ := connection.DB.DB()
		sqlDB.Close() // Cierra la conexión al final
	}()

	// Crear un router y configurar rutas
	router := gin.Default()
	mapa.ConfigurarRutas(router)

	// Iniciar el servidor
	log.Println("Servidor escuchando en :8080")
	router.Run(":8080")
}
