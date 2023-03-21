package routes

import (
	"petlover/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("pets", controllers.GetPetController)
	e.POST("pets", controllers.CreatePetController)
}
