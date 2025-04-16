package handdlers

import (
	"MegaModa/DB"
	"MegaModa/modelos"
	"MegaModa/services"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ObtenerDatosDeCookie(c *gin.Context) (uint, uint, error) {
	// Obtener el token JWT de la cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		return 0, 0, err
	}

	// Parsear el token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("clave_secreta"), nil
	})
	if err != nil || !token.Valid {
		return 0, 0, err
	}

	// Obtener los claims del token
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, 0, errors.New("token inválido")
	}

	return claims.IDLocal, claims.UsuarioID, nil
}

func obtenerTarjetas(c *gin.Context) ([]modelos.Tarjeta, error) {
	// Crear una instancia del servicio LocalTarjetaService
	tarjetaLocalService := services.NewLocalTarjetaService(DB.GlobalDB)

	// Obtener el ID del local desde los claims del token JWT
	localID := c.MustGet("id_local").(uint)

	// Obtener las tarjetas asociadas al local
	tarjetas, err := tarjetaLocalService.ObtenerTarjetasPorLocal(localID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener las tarjetas del local: %v", err)
	}

	return tarjetas, nil
}
func ServeIndexPage(c *gin.Context) {
	// Obtener el ID del local y el ID del usuario desde la cookie
	localID, usuarioID, err := ObtenerDatosDeCookie(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
		return
	}

	// Obtener las tarjetas asociadas al local
	tarjetas, err := obtenerTarjetas(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Imprimir los datos obtenidos de la base de datos
	log.Println("Tarjetas obtenidas de la base de datos:")
	for _, tarjeta := range tarjetas {
		log.Printf("ID: %d, Tipo: %s, Nombre: %s\n", tarjeta.ID, tarjeta.Tipo, tarjeta.Nombre)
	}
	//Obtener el nombre del local
	localService := services.NewLocalService(DB.GlobalDB)
	local, err := localService.ObtenerLocal(int(localID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el nombre del local"})
		return
	}
	// Renderizar la página HTML con los datos
	c.HTML(http.StatusOK, "index.html", gin.H{
		"local_id":     localID,
		"local_nombre": local.Nombre,
		"usuario_id":   usuarioID,
		"tarjetas":     tarjetas, // Pasar el array de tarjetas al HTML
	})
}
