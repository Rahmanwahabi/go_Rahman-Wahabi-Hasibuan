package database

import (
	"remedial/config"
	"remedial/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *models.Item) (*models.Item, error)
	GetItemByName(name string) ([]models.Item, error)
	GetAllItems() ([]models.Item, error)
	GetItemByID(id string) (*models.Item, error)
	UpdateItem(item *models.Item) (*models.Item, error)
	DeleteItem(item *models.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (i *itemRepository) CreateItem(item *models.Item) (*models.Item, error) {
	err := config.DB.Create(item).Error
	if err != nil {
		return nil, err
	}
	err = config.DB.Preload("Category").First(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemRepository) GetItemByName(name string) ([]models.Item, error) {
	var items []models.Item

	err := config.DB.Where("name LIKE ? ", "%"+name+"%").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := config.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemRepository) GetItemByID(id string) (*models.Item, error) {
	var item models.Item
	err := config.DB.Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (i *itemRepository) UpdateItem(item *models.Item) (*models.Item, error) {
	err := config.DB.Save(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemRepository) DeleteItem(item *models.Item) error {
	err := config.DB.Delete(&item).Error
	if err != nil {
		return err
	}
	return nil
}
