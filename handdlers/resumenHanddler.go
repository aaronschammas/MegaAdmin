package handdlers

import (
	"MegaModa/DB"
	"MegaModa/modelos"
	"MegaModa/services"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// coso que cree para tapar el hueco de la relacion entre turno y chicas ta medio rara pero funciona
type TurnoConChicas struct {
	modelos.Turno
	Chicas []modelos.Chica `json:"chicas"`
}

// datos en default de la primera solicitud
func obtenerDatosDefault() (map[string]interface{}, error) {
	// Valores por defecto
	fechaDesde := time.Now().Format("2006-01-02")
	fechaHasta := fechaDesde // Solo la fecha actual
	turno := "Mañana"
	//idLocal := uint(0)

	// Llamar a la función para obtener los datos filtrados
	return obtenerDatosFiltrados(uint(0), fechaDesde, fechaHasta, turno)
}

func FiltrarResumen(c *gin.Context) {
	// Obtener los parámetros de la URL
	fechaDesde := c.Query("fechaDesde")
	fechaHasta := c.Query("fechaHasta")
	turno := c.Query("turno")
	idLocal := c.Query("idLocal")

	// Si no se especifican filtros, usar valores por defecto
	if fechaDesde == "" {
		fechaDesde = time.Now().Format("2006-01-02")
	}
	if fechaHasta == "" {
		fechaHasta = fechaDesde
	}
	if turno == "" {
		turno = "Mañana"
	}
	if idLocal == "" {
		idLocal = "0"
	}

	// Llamar a la función que obtiene los datos del resumen
	resultado, err := obtenerDatosFiltrados(uint(0), fechaDesde, fechaHasta, turno)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos: " + err.Error()})
		return
	}

	// Enviar los datos en formato JSON al cliente
	c.JSON(http.StatusOK, resultado)
}

func obtenerDatosFiltrados(idLocal uint, fechaDesde, fechaHasta, turnoParam string) (map[string]interface{}, error) {
	// Inicializamos los servicios
	turnoService := services.NewTurnoService(DB.GlobalDB)
	localService := services.NewLocalService(DB.GlobalDB)

	// Obtenemos todos los locales
	locales, err := localService.ObtenerLocales()
	if err != nil {
		return nil, fmt.Errorf("error al obtener locales: %v", err)
	}

	// Obtenemos los turnos filtrados con chicas
	turnosConChicas, err := turnoService.ObtenerTurnosConChicasYFechas(idLocal, fechaDesde, fechaHasta, turnoParam)
	if err != nil {
		return nil, fmt.Errorf("error al obtener turnos con chicas: %v", err)
	}

	// Inicializamos variables para los totales y datos
	var totalEfectivo, totalCredito, totalDebito, totalEgresos float64
	var datos []struct {
		ID             uint
		Nombre         string
		TotalVentas    float64
		Efectivo       float64
		TarjetaCredito float64
		TarjetaDebito  float64
		TotalEgresos   float64
		Chicas         string
	}

	// Procesamos los turnos filtrados
	for _, turnoConChicas := range turnosConChicas {
		// Actualizamos los totales
		totalEfectivo += turnoConChicas.Efectivo
		totalCredito += turnoConChicas.TarjetaCredito
		totalDebito += turnoConChicas.TarjetaDebito
		totalEgresos += turnoConChicas.TotalEgresos

		// Concatenamos los nombres de las chicas
		var nombresChicas []string
		for _, chica := range turnoConChicas.Chicas {
			nombresChicas = append(nombresChicas, chica.Nombre)
		}

		// Obtenemos el nombre del local correspondiente
		var localNombre string
		for _, local := range locales {
			if local.ID == turnoConChicas.IDLocal {
				localNombre = local.Nombre
				break
			}
		}

		// Agregamos los datos al resultado
		datos = append(datos, struct {
			ID             uint
			Nombre         string
			TotalVentas    float64
			Efectivo       float64
			TarjetaCredito float64
			TarjetaDebito  float64
			TotalEgresos   float64
			Chicas         string
		}{
			ID:             turnoConChicas.IDLocal,
			Nombre:         localNombre,
			TotalVentas:    turnoConChicas.TotalVentas,
			Efectivo:       turnoConChicas.Efectivo,
			TarjetaCredito: turnoConChicas.TarjetaCredito,
			TarjetaDebito:  turnoConChicas.TarjetaDebito,
			TotalEgresos:   turnoConChicas.TotalEgresos,
			Chicas:         strings.Join(nombresChicas, ", "),
		})
	}

	// Creamos un mapa con los datos obtenidos
	resultado := map[string]interface{}{
		"total_efectivo": totalEfectivo,
		"total_credito":  totalCredito,
		"total_debito":   totalDebito,
		"total_egresos":  totalEgresos,
		"datos":          datos,
	}

	return resultado, nil
}

func MostrarResumen(c *gin.Context) {
	resultado, err := obtenerDatosDefault()
	if err != nil {
		// Si ocurre un error, se envía el mensaje al cliente
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ocurrió un error al cargar los datos predeterminados: " + err.Error()})
		return
	}

	// Log para verificar el contenido de resultado antes de renderizar
	log.Printf("Datos enviados a la plantilla: %+v", resultado)

	// Renderizamos la plantilla HTML con los datos obtenidos
	c.HTML(http.StatusOK, "resumenes.html", resultado)
}
