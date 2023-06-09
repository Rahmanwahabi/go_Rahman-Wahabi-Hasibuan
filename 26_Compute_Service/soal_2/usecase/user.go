package usecase

import (
	"tugas/models"
	"tugas/repository/database"
)

func CreateUser(user *models.User) error {
	return database.CreateUser(user)
}

func GetUserByID(id uint) (*models.User, error) {
	return database.GetUserByID(id)
}

func GetAllUsers() ([]models.User, error) {
	return database.GetAllUsers()
}

func UpdateUser(user *models.User) error {
	return database.UpdateUser(user)
}

func DeleteUser(user *models.User) error {
	return database.DeleteUser(user)
}
