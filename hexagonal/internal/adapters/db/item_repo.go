package repository

import (
	"database/sql"

	"github.com/osalomon/market-api/internal/domain/item"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Save(item item.Item) error {
	return nil
}
