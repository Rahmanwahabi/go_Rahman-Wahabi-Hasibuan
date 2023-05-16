package config

import (
	"fmt"

	"go_Rahman_Wahabi_Hasibuan/23_Unit_Testing/Praktikum/RESTful_API_Testing/models" // Tambahkan import ini

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_USER = "root"
const DB_PASS = "123Rahman"
const DB_HOST = "localhost"
const DB_PORT = "3306"
const DB_NAME = "unt"

const DB_USER_TEST = "root"
const DB_PASS_TEST = "123Rahman."
const DB_HOST_TEST = "localhost"
const DB_PORT_TEST = "3306"
const DB_NAME_TEST = "unt_test"

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&models.User{})
}

func InitDBTest() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrateTest()
}

func InitMigrateTest() {
	DB.Migrator().DropTable(&models.User{})
	DB.AutoMigrate(&models.User{})
}
