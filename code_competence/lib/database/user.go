package database

import (
	"remedial/config"
	"remedial/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
