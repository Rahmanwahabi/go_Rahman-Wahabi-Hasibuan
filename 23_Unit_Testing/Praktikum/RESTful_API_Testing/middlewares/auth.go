package middlewares

import (
	"fmt"
	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/config" // Tambahkan import ini
	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/models" // Tambahkan import ini

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var user models.User
	fmt.Println(email, password)
	if err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}
