package handdlers

import (
	"MegaModa/DB"
	"MegaModa/modelos"
	"MegaModa/services"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ObtenerLocales función para obtener los locales de un usuario

func ServeLocalesPage(c *gin.Context) {
	// Obtener el token JWT de la cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
		return
	}

	// Verificar el token JWT
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("clave_secreta"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Obtener los claims del token JWT
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Mostrar el ID del usuario en la consola
	log.Println("Buscando permisos para el usuario con ID:", claims.UsuarioID)

	// Obtener los permisos del usuario
	permisoService := services.NewPermisoService(DB.GlobalDB)
	permisos, err := permisoService.ObtenerPermisosPorIDUsuario(claims.UsuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los permisos del usuario"})
		return
	}

	// Obtener los locales asociados a cada permiso
	localService := services.NewLocalService(DB.GlobalDB)
	var locales []modelos.Local
	for _, permiso := range permisos {
		local, err := localService.ObtenerLocal(int(permiso.IDLocal))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los locales del usuario"})
			return
		}
		locales = append(locales, *local)
	}

	// Renderizar la página locales.html con los locales
	c.HTML(http.StatusOK, "locales.html", gin.H{
		"title":   "Locales",
		"locales": locales,
	})
}

func CompletarCookie(c *gin.Context) {
	var data struct {
		IDLocal string `json:"id_local"`
	}

	// Vincular datos enviados
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println("Error al vincular datos:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Convertir el IDLocal a uint
	idLocal, err := strconv.Atoi(data.IDLocal)
	if err != nil {
		log.Println("Error al convertir el IDLocal:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "IDLocal inválido"})
		return
	}

	// Leer el token existente de la cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		log.Println("Error al leer el token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no encontrado"})
		return
	}

	// Parsear el token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("clave_secreta"), nil
	})
	if err != nil || !token.Valid {
		log.Println("Error al parsear el token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Obtener los claims del token
	claims, ok := token.Claims.(*Claims)
	if !ok {
		log.Println("Error al obtener los claims:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		return
	}

	// Actualizar los claims con el nuevo IDLocal
	claims.IDLocal = uint(idLocal)

	// Crear un nuevo token con los claims actualizados
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := newToken.SignedString([]byte("clave_secreta"))
	if err != nil {
		log.Println("Error al crear el nuevo token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el nuevo token"})
		return
	}

	// Actualizar la cookie
	c.SetCookie("token", newTokenString, 3600*24*3, "/", "", false, true)
	log.Println("Local guardado exitosamente")
	c.JSON(http.StatusOK, gin.H{"message": "Local guardado exitosamente"})

}
