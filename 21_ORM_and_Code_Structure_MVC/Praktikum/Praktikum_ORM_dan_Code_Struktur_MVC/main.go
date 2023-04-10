package main

import (
	"go_Rahman_Wahabi_Hasibuan/21_ORM_and_Code_Structure_MVC/Praktikum/Praktikum_ORM_dan_Code_Struktur_MVC/controllers"
	"go_Rahman_Wahabi_Hasibuan/21_ORM_and_Code_Structure_MVC/Praktikum/Praktikum_ORM_dan_Code_Struktur_MVC/models"

	"github.com/labstack/echo/v4"
)

//ini sudah termasuk Soal Prioritas 1, 2 dan Ekplorasi

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

	e.GET("/books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBook)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.DELETE("/books/:id", controllers.DeleteBook)

	e.GET("/blogs", controllers.GetAllBlogs)
	e.GET("/blogs/:id", controllers.GetBlog)
	e.POST("/blogs", controllers.CreateBlog)
	e.PUT("/blogs/:id", controllers.UpdateBlog)
	e.DELETE("/blogs/:id", controllers.DeleteBlog)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
