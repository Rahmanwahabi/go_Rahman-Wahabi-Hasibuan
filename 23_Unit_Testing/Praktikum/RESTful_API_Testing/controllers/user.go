package controllers

import (
	"net/http"

	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/config"
	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/middlewares"
	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/models"

	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	Message string
	Data    []models.User
}

func Addition(a, b int) int {
	result := a + b
	return result
}

func GetUserController(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {

	user := models.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail login user",
			"status":  err.Error(),
		})
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   user,
	})
}
