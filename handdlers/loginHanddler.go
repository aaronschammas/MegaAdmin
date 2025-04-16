package handdlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler para mostrar la página de login
func ServeLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", nil)
}
