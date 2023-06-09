package database

import (
	"tugas/config"
	"tugas/models"
)

func CreateUser(user *models.User) error {
	db := config.GetDB()
	return db.Create(&user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]models.User, error) {
	db := config.GetDB()

	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(user *models.User) error {
	db := config.GetDB()
	return db.Save(&user).Error
}

func DeleteUser(user *models.User) error {
	db := config.GetDB()
	return db.Delete(&user).Error
}
