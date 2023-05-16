package main

import (
	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/config"

	routes "go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/route"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":8080")
}
