package item

import (
	"github.com/osalomon/market-api/internal/domain/item"
)

type ItemUseCase interface {
	CreateItem(name string) error
}

type itemUseCase struct {
	repo item.Repository
}

func NewItemUseCase(repo item.Repository) *itemUseCase {
	return &itemUseCase{repo: repo}
}

func (uc *itemUseCase) CreateItem(name string) error {
	i := item.Item{Name: name}
	return uc.repo.Save(i)
}
