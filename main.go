package main

import (
	"petlover/config"
	"petlover/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	routes.InitRoutes(e)
	e.Start(":8000")
}
