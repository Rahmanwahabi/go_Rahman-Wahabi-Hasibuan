package main

import (
	"remedial/config"
	"remedial/middlewares"
	"remedial/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middlewares.Logger()

	routes.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
