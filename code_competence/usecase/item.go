package usecase

import (
	"remedial/lib/database"
	"remedial/models"
	"remedial/models/payload"

	"github.com/google/uuid"
)

type ItemUseCase interface {
	CreateItem(req payload.CreateItemRequest) (*models.Item, error)
	GetItemByName(name string) ([]models.Item, error)
	GetAllItems() ([]models.Item, error)
	GetItemByID(id string) (*models.Item, error)
	UpdateItemByID(id string, item models.Item) (*models.Item, error)
	DeleteItemByID(id string) error
}

type itemUseCase struct {
	itemRepository database.ItemRepository
}

func NewItemUseCase(itemRepository database.ItemRepository) *itemUseCase {
	return &itemUseCase{itemRepository: itemRepository}
}

func (i *itemUseCase) CreateItem(req payload.CreateItemRequest) (*models.Item, error) {
	id, _ := uuid.NewRandom()
	item := &models.Item{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	item, err := i.itemRepository.CreateItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemUseCase) GetItemByName(name string) ([]models.Item, error) {
	item, err := i.itemRepository.GetItemByName(name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemUseCase) GetAllItems() ([]models.Item, error) {
	items, err := i.itemRepository.GetAllItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemUseCase) GetItemByID(id string) (*models.Item, error) {
	item, err := i.itemRepository.GetItemByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemUseCase) UpdateItemByID(id string, item models.Item) (*models.Item, error) {
	items, err := i.itemRepository.GetItemByID(id)
	if err != nil {
		return nil, err
	}
	items.Name = item.Name
	items.Description = item.Description
	items.Price = item.Price
	items.Stock = item.Stock
	items.CategoryID = item.CategoryID
	newItem, err := i.itemRepository.UpdateItem(items)
	if err != nil {
		return nil, err
	}
	return newItem, nil

}

func (i *itemUseCase) DeleteItemByID(id string) error {
	items, err := i.itemRepository.GetItemByID(id)
	if err != nil {
		return err
	}

	err = i.itemRepository.DeleteItem(items)
	if err != nil {
		return err
	}
	return nil
}
