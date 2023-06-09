package config

import (
	"fmt"
	"tugas/models"

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
		DB_Username: "admin",
		DB_Password: "Killerqueen2002",
		DB_Port:     "3307",
		DB_Host:     "mariadb1.c8gn4dschw0y.ap-southeast-2.rds.amazonaws.com",
		DB_Name:     "awsdatabase",
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
	DB.AutoMigrate(&models.User{})
}
