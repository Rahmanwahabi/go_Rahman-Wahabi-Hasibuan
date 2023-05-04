package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	"go_Rahman_Wahabi_Hasibuan/22_Middleware/middleware"
	"go_Rahman_Wahabi_Hasibuan/22_Middleware/models"
)

type UserController struct {
	DB *gorm.DB
}

type BookController struct {
	DB *gorm.DB
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users := []models.User{}

	// Membuat koneksi ke database
	//models.InitDB()

	// Mengambil semua data user dari database
	uc.DB.Find(&users)

	// Menutup koneksi ke database
	//defer models.DB.Close()

	return c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}

	// Membuat koneksi ke database
	//models.InitDB()

	// Mengambil data user berdasarkan ID dari database
	uc.DB.First(&user, id)

	// Menutup koneksi ke database
	//defer models.DB.Close()

	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateNewUser(c echo.Context) error {
	user := models.User{}

	// Binding request body ke struct User
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Membuat koneksi ke database
	//models.InitDB()

	// Menambahkan user baru ke database
	uc.DB.Create(&user)

	// Menutup koneksi ke database
	//defer models.DB.Close()

	return c.JSON(http.StatusCreated, user)
}

func (uc *UserController) UpdateUserByID(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}

	// Binding request body ke struct User
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Membuat koneksi ke database
	//models.InitDB()

	// Mengupdate data user berdasarkan ID di database
	uc.DB.Model(&user).Where("id = ?", id).Updates(models.User{Name: user.Name, Email: user.Email, Password: user.Password})

	// Menutup koneksi ke database
	//defer models.DB.Close()

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUserByID(c echo.Context) error {
	id := c.Param("id")
	user := models.User{}

	// Membuat koneksi ke database
	//models.InitDB()

	// Menghapus data user berdasarkan ID di database
	uc.DB.Where("id = ?", id).Delete(&user)

	// Menutup koneksi ke database
	//defer models.DB.Close()

	return c.NoContent(http.StatusNoContent)
}

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	// Declare and initialize the config variable
	config := struct {
		DB *gorm.DB
	}{DB: models.DB}

	uc := UserController{DB: config.DB}
	err := uc.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{
		ID:    uint64(user.ID),
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "",
		"user":    userResponse,
	})
}

func (bc *BookController) GetAllBooks(c echo.Context) error {
	books := []models.Book{}

	bc.DB.Find(&books)

	return c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}

	bc.DB.First(&book, id)

	if book.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	return c.JSON(http.StatusOK, book)
}

func (bc *BookController) CreateNewBook(c echo.Context) error {
	book := models.Book{}

	if err := c.Bind(&book); err != nil {
		return err
	}

	bc.DB.Create(&book)

	return c.JSON(http.StatusCreated, book)
}

func (bc *BookController) UpdateBookByID(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}

	if err := c.Bind(&book); err != nil {
		return err
	}

	bc.DB.Model(&book).Where("id = ?", id).Updates(models.Book{Name: book.Name, Email: book.Email, Password: book.Password})

	return c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBookByID(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}

	bc.DB.Where("id = ?", id).Delete(&book)

	return c.NoContent(http.StatusNoContent)
}

func LoginBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	// Declare and initialize the config variable
	config := struct {
		DB *gorm.DB
	}{DB: models.DB}

	uc := UserController{DB: config.DB}
	err := uc.DB.Where("email = ? AND password = ?", book.Email, book.Password).First(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateBookToken(int(book.ID), book.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{
		ID:    uint64(book.ID),
		Name:  book.Name,
		Email: book.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "",
		"user":    userResponse,
	})
}
