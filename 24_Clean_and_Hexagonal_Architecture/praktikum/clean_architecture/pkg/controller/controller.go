package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userControllerPkg "code/internal/user/controller"
	userRepositoryPkg "code/internal/user/repository"
	userServicePkg "code/internal/user/service"
)

func InitControllers(e *echo.Echo, db *gorm.DB) {
	userRepository := userRepositoryPkg.NewUserRepositoryImpl(db)
	userService := userServicePkg.NewUserServiceImpl(userRepository)
	userController := userControllerPkg.NewUserController(userService)
	userController.InitRoutes(e)
}
