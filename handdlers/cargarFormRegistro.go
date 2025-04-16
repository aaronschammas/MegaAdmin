package handdlers

import (
	"MegaModa/DB"
	"MegaModa/modelos"
	"MegaModa/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// cargar pagina de registro
func ServeRegistro(c *gin.Context) {
	c.HTML(http.StatusOK, "registro.html", nil)
}

// Registrar usuarios
func CargarUsuario(c *gin.Context) {
	usuarioService := services.NewUsuarioService(DB.GlobalDB)

	var input struct {
		Nombre     string `form:"nombre" binding:"required"`
		Contraseña string `form:"contraseña" binding:"required"`
		IdRol      string `form:"id_rol" binding:"required"`
	}

	// Intentar hacer el binding
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convertir id_rol de string a uint
	idRol, err := strconv.ParseUint(input.IdRol, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de rol inválido"})
		return
	}

	// Crear usuario con los datos correctos
	nuevoUsuario := modelos.Usuario{
		Nombre:     input.Nombre,
		Contraseña: input.Contraseña,
		Id_rol:     uint(idRol),
	}

	// Guardar en la base de datos
	if err := usuarioService.CrearUsuario(&nuevoUsuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario registrado correctamente"})
}
