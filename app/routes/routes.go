package routes

import (
	m "hungry-baby/app/middleware"
	"hungry-baby/controllers/auth"
	"hungry-baby/controllers/city"
	"hungry-baby/controllers/country"
	"hungry-baby/controllers/file"
	"hungry-baby/controllers/province"
	"hungry-baby/controllers/user"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	FileController     file.FileController
	CountryController  country.CountryController
	ProvinceController province.ProvinceController
	CityController     city.CityController
	UserController     user.UserController
	AuthController     auth.AuthController
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

	province := v1.Group("/province")
	province.Use(middleware.JWTWithConfig(c.JWTMiddleware))
	province.Use(m.LoadClaims(c.JWTMiddleware))
	province.GET("/all", c.ProvinceController.FindAll)
	province.GET("", c.ProvinceController.Find)
	province.GET("/:id", c.ProvinceController.FindByID)
	province.POST("", c.ProvinceController.Store)
	province.PUT("/:id", c.ProvinceController.Update)
	province.DELETE("/:id", c.ProvinceController.Delete)

	city := v1.Group("/city")
	city.Use(middleware.JWTWithConfig(c.JWTMiddleware))
	city.Use(m.LoadClaims(c.JWTMiddleware))
	city.GET("/all", c.CityController.FindAll)
	city.GET("", c.CityController.Find)
	city.GET("/:id", c.CityController.FindByID)
	city.POST("", c.CityController.Store)
	city.PUT("/:id", c.CityController.Update)
	city.DELETE("/:id", c.CityController.Delete)

	user := v1.Group("/user")
	user.Use(middleware.JWTWithConfig(c.JWTMiddleware))
	user.Use(m.LoadClaims(c.JWTMiddleware))
	user.GET("/all", c.UserController.FindAll)
	user.GET("", c.UserController.Find)
	user.GET("/:id", c.UserController.FindByID)
	user.POST("", c.UserController.Store)
	user.PUT("/:id", c.UserController.Update)
	user.DELETE("/:id", c.UserController.Delete)

	auth := v1.Group("/auth")
	auth.GET("/loginUrl", c.AuthController.GetGoogleLoginURL)
	auth.GET("/google", c.AuthController.VerifyGoogleCode)
}
