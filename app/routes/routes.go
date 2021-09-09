package routes

import (
	m "hungry-baby/app/middleware"
	"hungry-baby/controllers/country"
	"hungry-baby/controllers/file"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	FileController    file.FileController
	CountryController country.CountryController
}

func (c *ControllerList) RouteRegister(e *echo.Echo) {
	v1 := e.Group("/v1")

	file := v1.Group("/file")
	file.Use(middleware.JWTWithConfig(c.JWTMiddleware))
	file.Use(m.LoadClaims(c.JWTMiddleware))
	file.GET("/:id", c.FileController.FindByID)
	file.POST("", c.FileController.Upload)
	file.DELETE("/:id", c.FileController.Delete)

	country := v1.Group("/country")
	country.Use(middleware.JWTWithConfig(c.JWTMiddleware))
	country.Use(m.LoadClaims(c.JWTMiddleware))
	country.GET("/all", c.CountryController.FindAll)
	country.GET("", c.CountryController.Find)
	country.GET("/:id", c.CountryController.FindByID)
	country.POST("", c.CountryController.Store)
	country.PUT("/:id", c.CountryController.Update)
	country.DELETE("/:id", c.CountryController.Delete)
}
