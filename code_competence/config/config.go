package config

import (
	"fmt"
	"remedial/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func GetDB() *gorm.DB {
	return DB
}

func InitDB() *gorm.DB {
	config := Config{
		DB_Username: "root",
		DB_Password: "123Rahman",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "warung",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Item{})
}
