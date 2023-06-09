package routes

import (
	"tugas/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", controllers.CreateUser)
		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/users/:id", controllers.GetUserByID)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}

	return r
}
