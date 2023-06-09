package routes

import (
	"net/http"
	"remedial/constant"
	"remedial/controllers"
	"remedial/lib/database"
	"remedial/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := database.NewUserRepository(db)
	itemRepository := database.NewItemRepository(db)
	categoryRepository := database.NewCategoryRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	itemUsecase := usecase.NewItemUseCase(itemRepository)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)

	userController := controllers.NewUserController(userUsecase)
	itemController := controllers.NewItemController(itemUsecase)
	categoryController := controllers.NewCategoryController(categoryUsecase)

	e.Validator = &customValidator{validator: validator.New()}

	e.POST("/register", userController.RegisterUserController)
	e.POST("/login", userController.LoginUserController)

	items := e.Group("/items", middleware.JWT([]byte(constant.SECRET_JWT)))
	items.POST("", itemController.CreateItemController)
	items.GET("", itemController.GetAllItemsController)
	items.GET("/id/:id", itemController.GetItemByIDController)
	items.PUT("/id/:id", itemController.UpdateItemByIDController)
	items.DELETE("/id/:id", itemController.DeleteItemByIDController)
	items.POST("/category", categoryController.CreateCategoryController)
	items.GET("/category/:category_id", categoryController.GetAllByCategoryIDController)
	items.GET("/name", itemController.GetItemByItemNameController)
}
