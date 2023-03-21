package main

import (
	"os"
	"petlover/config"
	"petlover/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	routes.InitRoutes(e)
	port := os.Getenv("PORT")
	if port == "" {
 		port = "8080"
	}
	e.Start(":" + port)

}
