package main

import (
	"github.com/labstack/echo/v4"

	"go_Rahman_Wahabi_Hasibuan/22_Middleware/constants"
	"go_Rahman_Wahabi_Hasibuan/22_Middleware/controllers"
	"go_Rahman_Wahabi_Hasibuan/22_Middleware/models"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Initialize database
	models.InitDB()

	// Run migrations
	models.InitialMigration()

	e.Use(middleware.Logger())

	// create new UserController instance and pass DB connection
	userController := controllers.UserController{DB: models.DB}
	bookController := controllers.BookController{DB: models.DB}
	// Routes
	e.POST("/login", controllers.LoginUserController) // Route untuk login
	e.POST("/loginbook", controllers.LoginBookController)
	// Routes yang memerlukan otentikasi JWT
	eApi := e.Group("/api")
	eApi.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eApi.GET("/users", userController.GetAllUsers)
	eApi.GET("/users/:id", userController.GetUserByID)
	eApi.POST("/users", userController.CreateNewUser)
	eApi.PUT("/users/:id", userController.UpdateUserByID)
	eApi.DELETE("/users/:id", userController.DeleteUserByID)

	eApi.GET("/books", bookController.GetAllBooks)
	eApi.GET("/books/:id", bookController.GetBookByID)
	eApi.POST("/books", bookController.CreateNewBook)
	eApi.PUT("/books/:id", bookController.UpdateBookByID)
	eApi.DELETE("/books/:id", bookController.DeleteBookByID)
	// Start server

	e.Logger.Fatal(e.Start(":8080"))
}
