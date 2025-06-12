package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/osalomon/meli-items-v2/internal/entity"
)

type ItemRepository interface {
	GetItems() ([]entity.Item, error)
	GetItemByID(id uint) (entity.Item, error)
	AddItem(item *entity.Item) error
}

type mysqlItemRepository struct {
	conn *sqlx.DB
}

func NewMySQLItemRepository(db *sqlx.DB) *mysqlItemRepository {
	return &mysqlItemRepository{
		conn: db,
	}
}

func (r *mysqlItemRepository) GetItemByID(id uint) (entity.Item, error) {
	item := entity.Item{}
	err := r.conn.Get(&item, "SELECT id, code, title, description, price, stock, status, created_at, updated_at FROM items WHERE id = ?", id)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *mysqlItemRepository) AddItem(item *entity.Item) error {
	// TODO: implement
	return nil
}

func (r *mysqlItemRepository) GetItems() ([]entity.Item, error) {
	// TODO: implement
	return nil, nil
}
