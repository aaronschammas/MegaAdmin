package handdlers

import (
	"MegaModa/DB"
	"MegaModa/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServeTurnosNullPage función para obtener turnos con turno NULL
func ServeTurnosNullPage(c *gin.Context) {
	turnoService := services.NewTurnoService(DB.GlobalDB)
	turnos, err := turnoService.ObtenerTurnosNull()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "turnos-null.html", gin.H{
		"turnos": turnos,
	})
}

// Crear la variable globlal de turnos para poder usarla en la otra funcion

func ActualizarTurnos(c *gin.Context) {
	var turnosModificados []services.TurnosModificados // Usar el tipo correcto

	if err := json.NewDecoder(c.Request.Body).Decode(&turnosModificados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar datos: " + err.Error()})
		return
	}

	// Llamar al servicio para actualizar los turnos en la DB
	turnoService := services.NewTurnoService(DB.GlobalDB)
	err := turnoService.ActualizarTurnos(turnosModificados)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar turnos: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Turnos actualizados correctamente"})
}
