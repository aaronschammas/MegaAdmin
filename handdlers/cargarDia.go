package handdlers

import (
	"MegaModa/DB"
	"MegaModa/modelos"
	"MegaModa/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Estructura para recibir los datos del turno y las tarjetas
type CargarTurno struct {
	Turno          modelos.Turno            `json:"turno"`
	Chicas         []modelos.Chica          `json:"chicas"`
	ConceptoEgreso []modelos.ConceptoEgreso `json:"ConceptoEgreso"` // Cambia esto
	Vouchers       []modelos.Voucher        `json:"vouchers"`
	TotalVentas    float64                  `json:"total_ventas"` // Nuevo campo
}

func CrearTurno(c *gin.Context) {
	var request CargarTurno

	// Bindear los datos JSON a la estructura
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error al bindear JSON: %v", err)})
		return
	}

	// Validar si el array de ConceptosEgreso está vacío
	if len(request.ConceptoEgreso) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El array de conceptos de egreso está vacío"})
		fmt.Println("Toy vacio")
		return
	}

	// Calcular el total de ventas
	totalVentas := request.Turno.Efectivo + request.Turno.TarjetaCredito + request.Turno.TarjetaDebito + request.Turno.Financiera - request.Turno.TotalEgresos

	// Asignar el valor de total de ventas al campo correspondiente
	request.Turno.TotalVentas = totalVentas

	// Crear el turno
	turnoID, err := crearTurno(&request.Turno, DB.GlobalDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al crear el turno: %v", err)})
		return
	}

	// Cargar los vouchers
	if err := cargarVouchers(request.Vouchers, turnoID, DB.GlobalDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al cargar los vouchers: %v", err)})
		return
	}

	// Cargar las chicas
	if err := cargarChicas(request.Chicas, turnoID, DB.GlobalDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al cargar las chicas: %v", err)})
		return
	}

	// Cargar los conceptos de egreso
	if err := cargarConceptosEgreso(request.ConceptoEgreso, turnoID, DB.GlobalDB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al cargar los conceptos de egreso: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datos cargados con éxito"})
}

// Función para crear el turno
func crearTurno(turno *modelos.Turno, db *gorm.DB) (uint, error) {
	turnoService := services.NewTurnoService(db)
	if err := turnoService.CrearTurno(turno); err != nil {
		return 0, fmt.Errorf("error al guardar el turno: %v", err)
	}

	return turno.ID, nil
}

// Función para cargar los vouchers
func cargarVouchers(vouchers []modelos.Voucher, turnoID uint, db *gorm.DB) error {
	if len(vouchers) == 0 {
		return fmt.Errorf("la lista de vouchers está vacía")
	}

	voucherService := services.NewVoucherService(db)
	for i, voucher := range vouchers {
		voucher.IDTurno = turnoID

		// Mostrar los datos del voucher antes de guardarlo
		fmt.Printf("Voucher a guardar (posición %d): %+v\n", i, voucher)

		if err := voucherService.CrearVoucher(&voucher); err != nil {
			return fmt.Errorf("error al guardar el voucher en la posición %d: %v", i, err)
		}
	}
	return nil
}

// Función para cargar las chicas
func cargarChicas(chicas []modelos.Chica, turnoID uint, db *gorm.DB) error {
	chicaService := services.NewChicaService(db)
	for _, chica := range chicas {
		chica.IDTurno = turnoID
		if err := chicaService.CrearChica(&chica); err != nil {
			return fmt.Errorf("error al guardar la chica: %v", err)
		}
	}
	return nil
}

// Función para cargar los concept os de egreso
func cargarConceptosEgreso(conceptos []modelos.ConceptoEgreso, turnoID uint, db *gorm.DB) error {
	conceptoService := services.NewConceptoEgresoService(db)
	for _, concepto := range conceptos {
		concepto.IDTurno = turnoID
		if err := conceptoService.CrearConceptoEgreso(&concepto); err != nil {
			return fmt.Errorf("error al guardar el concepto de egreso: %v", err)
		}
	}
	return nil
}
