package repository

import "github.com/teamcubation/neocamp-meli/testing/domain"

//go:generate mockery --name ItemRepository --output=../mocks --outpkg=mocks
type ItemRepository interface {
	// SaveItem(name string, stock int) error
	SaveItem(item domain.Item) error
	GetItemByID(itemID uint) error
}
