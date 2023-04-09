package main

import (
	"rahman/21_ORM_and_Code_Structure_MVC/Praktikum/controllers"
	"rahman/21_ORM_and_Code_Structure_MVC/Praktikum/models"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize database
	models.InitDB()

	// Run migrations
	models.InitialMigration()

	// Initialize Echo instance
	e := echo.New()

	// Routes for User CRUD Operations
	e.GET("/users", controllers.GetAllUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
