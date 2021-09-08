package routes

import (
	"hungry-baby/controllers/category"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	CategoryController category.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	category := e.Group("category")
	category.GET("/list", cl.CategoryController.GetAll, middleware.JWTWithConfig(cl.JWTMiddleware))
}
