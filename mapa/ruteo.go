package mapa

import (
	"MegaModa/handdlers" // Importar el package handdlers

	"github.com/gin-gonic/gin"
)

// ConfigurarRutas configura las rutas del servidor
func ConfigurarRutas(router *gin.Engine) {
	// Configuración global de plantillas y favicon
	router.LoadHTMLGlob("front/*.html")                      // Cargar plantillas HTML
	router.StaticFile("/favicon.ico", "./front/favicon.ico") // Favicon

	// Rutas para servir las páginas HTML
	router.GET("/", handdlers.MiddlewareAutenticacion(), handdlers.ServeIndexPage)

	router.GET("/login", handdlers.ServeLoginPage)
	router.GET("/locales", handdlers.ServeLocalesPage)

	router.POST("/completar-cookie", handdlers.CompletarCookie)
	router.POST("/cargar", handdlers.CrearTurno) //Cargar datos a la DB
	router.POST("/login", handdlers.Login)
	//registro de usuarios
	router.GET("/registro", handdlers.ServeRegistro)
	router.POST("/registro", handdlers.CargarUsuario)

	//rutas para los administradores.
	admin := router.Group("/admin")
	{

		admin.GET("", nil)

		admin.GET("/modificar-turnos", handdlers.ServeTurnosNullPage)
		admin.POST("/actualizar-turnos", handdlers.ActualizarTurnos)
		admin.POST("/registro")

		//rutas para manejar usuarios
		admin.GET("/usuarios")
		//rutas para los resumenes
		resumenes := router.Group("/resumenes")
		{
			// Ruta para el resumen predeterminado
			resumenes.GET("", handdlers.MiddlewareResumen(), handdlers.MostrarResumen)

			// Ruta para el resumen con filtros aplicados
			resumenes.GET("/filtrar", handdlers.MiddlewareResumen(), handdlers.FiltrarResumen)
		}
	}
	// Servir archivos estáticos (CSS)
	router.Static("/css", "./front/css")
	router.Static("/js", "./front/js") // Servir archivos js
}
