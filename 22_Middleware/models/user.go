package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

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
		DB_Name:     "mvc_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"Name"`
	Email    string `json:"email" form:"Email"`
	Password string `json:"password" form:"Password"`
}

type UserResponse struct {
	gorm.Model

	ID    uint64 `json:"id" form:"Name"`
	Name  string `json:"name" form:"Name"`
	Email string `json:"email" form:"Email"`
	Token string `json:"token" form:"Token"`
}

type Book struct {
	gorm.Model
	Name     string `json:"name" form:"Name"`
	Email    string `json:"email" form:"Email"`
	Password string `json:"password" form:"Password"`
}

type BookResponse struct {
	gorm.Model

	ID    uint64 `json:"id" form:"Name"`
	Name  string `json:"name" form:"Name"`
	Email string `json:"email" form:"Email"`
	Token string `json:"token" form:"Token"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{}, &Book{})
}
