package database

import (
	"remedial/config"
	"remedial/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(req *models.Category) (*models.Category, error)
	GetAllItemByCategoryID(id string) ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	err := config.DB.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryRepository) GetAllItemByCategoryID(id string) ([]models.Category, error) {
	var category []models.Category
	err := config.DB.Preload("Item").Where("id = ?", id).Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
