package controllers

import (
	"errors"
	"go_Rahman_Wahabi_Hasibuan/21_ORM_and_Code_Structure_MVC/Praktikum/Praktikum_ORM_dan_Code_Struktur_MVC/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

	// Mencari user berdasarkan ID
	user := models.User{}
	if err := models.DB.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		}
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Bind data yang dikirim dari client ke struct user
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Update data user ke dalam database
	if err := models.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Response hasil update user
	return c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
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

func GetAllBooks(c echo.Context) error {
	books := []models.Book{}
	models.DB.Find(&books)
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}
	models.DB.First(&book, id)
	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, "Book not found")
	}
	return c.JSON(http.StatusOK, book)
}

func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	models.DB.Create(&book)
	return c.JSON(http.StatusOK, book)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}
	if err := models.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		}
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}
	models.DB.First(&book, id)
	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, "Book not found")
	}
	models.DB.Delete(&book)
	return c.JSON(http.StatusOK, "Book deleted")
}

func GetAllBlogs(c echo.Context) error {
	blogs := []models.Blog{}
	models.DB.Find(&blogs)
	return c.JSON(http.StatusOK, blogs)
}

func GetBlog(c echo.Context) error {
	id := c.Param("id")
	blog := models.Blog{}
	models.DB.First(&blog, id)
	if blog.ID == 0 {
		return c.JSON(http.StatusNotFound, "blog not found")
	}
	return c.JSON(http.StatusOK, blog)
}

func CreateBlog(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)
	models.DB.Create(&blog)
	return c.JSON(http.StatusOK, blog)
}

func UpdateBlog(c echo.Context) error {
	id := c.Param("id")
	blog := models.Blog{}
	if err := models.DB.First(&blog, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		}
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := c.Bind(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := models.DB.Save(&blog).Error; err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, blog)
}

func DeleteBlog(c echo.Context) error {
	id := c.Param("id")
	blog := models.Blog{}
	models.DB.First(&blog, id)
	if blog.ID == 0 {
		return c.JSON(http.StatusNotFound, "Blog not found")
	}
	models.DB.Delete(&blog)
	return c.JSON(http.StatusOK, "Blog deleted")
}
