package handdlers

import (
	"MegaModa/DB"
	"MegaModa/modelos"
	"MegaModa/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims estructura para almacenar los datos del usuario
type Claims struct {
	UsuarioID uint   `json:"usuario_id"`
	Nombre    string `json:"nombre"`
	IDLocal   uint   `json:"id_local"`
	RolID     uint   `json:"rol_id"`
	jwt.StandardClaims
}

// Autenticar función para autenticar al usuario

func Autenticar(c *gin.Context) {
	var usuario modelos.Usuario

	// Obtener los datos del formulario de login
	var login struct {
		Usuario  string `json:"usuario"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al bindear JSON"})
		return
	}

	// Buscar al usuario en la base de datos
	usuarioService := services.NewUsuarioService(DB.GlobalDB)
	usuario, err = usuarioService.ObtenerUsuarioPorNombre(login.Usuario)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Verificar la contraseña
	if usuario.Contraseña != login.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		return
	}

	// Obtener el Rol ID del usuario
	rolID := usuario.Id_rol

	// Crear un token JWT con los claims del usuario
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UsuarioID: usuario.ID,
		Nombre:    usuario.Nombre,
		IDLocal:   0, // Establecer un valor vacío para el campo IDLocal
		RolID:     rolID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "MegaModa",
		},
	})

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString([]byte("clave_secreta"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al firmar el token"})
		return
	}

	// Crear la cookie con el token JWT
	c.SetCookie("token", tokenString, 3600*24*3, "/", "", false, true) // Válida por 72 horas

	c.Redirect(http.StatusFound, "/locales")
}

func MiddlewareAutenticacion() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Intentar obtener el token desde la cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			// Si no hay token, redirigir al login
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Parsear el token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("clave_secreta"), nil
		})
		if err != nil || !token.Valid {
			// Si el token es inválido, redirigir al login
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Obtener los claims del token
		claims, ok := token.Claims.(*Claims)
		if !ok {
			// Si los claims no son válidos, redirigir al login
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Verificar si el IDLocal está configurado
		if claims.IDLocal == 0 {
			// Si IDLocal es 0, redirigir al selector de locales
			c.Redirect(http.StatusFound, "/locales")
			c.Abort()
			return
		}

		// Si todo está bien, almacenar los claims en el contexto
		c.Set("usuario_id", claims.UsuarioID)
		c.Set("nombre", claims.Nombre)
		c.Set("id_local", claims.IDLocal)

		// Continuar con la solicitud
		c.Next()
	}
}

func MiddlewareResumen() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token JWT de la cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		// Parsear el token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("clave_secreta"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Obtener los claims del token
		claims, ok := token.Claims.(*Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Verificar si el Rol ID es 2
		if claims.RolID != 2 {
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso para acceder a esta página"})
			c.Abort()
			return
		}

		// Continuar con la solicitud
		c.Next()
	}
}
