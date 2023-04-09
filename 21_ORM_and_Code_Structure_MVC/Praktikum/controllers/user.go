package controllers

import (
	"net/http"
	"rahman/21_ORM_and_Code_Structure_MVC/Praktikum/models"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users := []models.User{}
	models.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	models.DB.First(&user, id)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	models.DB.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	models.DB.First(&user, id)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	c.Bind(&user)
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}
	models.DB.First(&user, id)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	models.DB.Delete(&user)
	return c.JSON(http.StatusOK, "User deleted")
}
