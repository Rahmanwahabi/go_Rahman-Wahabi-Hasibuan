package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "123Rahman",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := config.DB_Username + ":" + config.DB_Password + "@tcp(" + config.DB_Host + ":" + config.DB_Port + ")/" + config.DB_Name + "?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

// get all users
func GetUsersController(c *gin.Context) {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get user",
		"user":    user,
	})
}

// create user
func CreateUserController(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success create user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user User
	if err := DB.First(&user, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update user",
		"user":    user,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "success update user",
		"user":    user,
	})
}

func main() {
	r := gin.Default()

	r.GET("/users", GetUsersController)
	r.GET("/users/:id", GetUserController)
	r.POST("/users", CreateUserController)
	r.DELETE("/users/:id", DeleteUserController)
	r.PUT("/users/:id", UpdateUserController)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
