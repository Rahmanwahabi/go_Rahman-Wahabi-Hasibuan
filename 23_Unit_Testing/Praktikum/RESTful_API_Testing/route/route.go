package routes

import (
	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/constants"
	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/controllers"

	"github.com/labstack/echo/v4"
	echomid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)

	// eBasicAuth := e.Group("")
	// eBasicAuth.Use(echomid.BasicAuth(middleware.BasicAuthDB))
	// eBasicAuth.GET("/users", controllers.GetUserController)

	jwtAuth := e.Group("user")
	jwtAuth.Use(echomid.JWT([]byte(constants.SECRET_JWT)))
	jwtAuth.GET("", controllers.GetUserController)
	return e
}
