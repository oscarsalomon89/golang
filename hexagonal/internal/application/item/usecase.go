package item

import (
	"github.com/osalomon/market-api/internal/domain/item"
)

type ItemUseCase interface {
	CreateItem(item *item.Item) error
	GetItemByID(id int) (item.Item, error)
	GetItems() ([]item.Item, error)
	DeleteItem(id int) error
}

type itemUseCase struct {
	repo item.Repository
}

func NewItemUseCase(repo item.Repository) *itemUseCase {
	return &itemUseCase{repo: repo}
}

func (uc *itemUseCase) CreateItem(item *item.Item) error {
	return uc.repo.Save(item)
}

func (uc *itemUseCase) GetItemByID(id int) (item.Item, error) {
	return uc.repo.GetItemByID(id)
}

func (uc *itemUseCase) GetItems() ([]item.Item, error) {
	return uc.repo.GetItems()
}

func (uc *itemUseCase) DeleteItem(id int) error {
	return uc.repo.DeleteItem(id)
}
