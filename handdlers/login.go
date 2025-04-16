package handdlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Obtener los valores enviados desde el formulario
	usuario := c.PostForm("usuario")
	password := c.PostForm("password")

	// Validar que los campos no estén vacíos
	if usuario == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario y la contraseña son requeridos"})
		return
	}

	// Crear una estructura con los datos del formulario
	loginData := struct {
		Usuario  string `json:"usuario"`
		Password string `json:"password"`
	}{
		Usuario:  usuario,
		Password: password,
	}

	// Convertir la estructura en JSON
	jsonData, err := json.Marshal(loginData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los datos"})
		return
	}

	// Reemplazar el cuerpo de la solicitud con el JSON generado
	c.Request.Body = http.NoBody // Limpiar el cuerpo original
	c.Request.Body = io.NopCloser(bytes.NewReader(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	// Llamar a la función Autenticar
	Autenticar(c)
}
