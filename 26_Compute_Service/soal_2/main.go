package main

import (
	"tugas/config"
	"tugas/routes"
)

func main() {
	config.InitDB()

	r := routes.SetupRoutes()
	r.Run(":8080")
}
