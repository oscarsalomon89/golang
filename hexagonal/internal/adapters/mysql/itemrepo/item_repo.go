package itemrepo

import (
	"time"

	"github.com/osalomon/market-api/internal/domain/item"
	"gorm.io/gorm"
)

type ItemDB struct {
	ID          int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"size:100;not null"`
	Price       float64 `gorm:"not null"`
	Description string  `gorm:"type:text"`
	Stock       int
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (ItemDB) TableName() string {
	return "items"
}

type UsersDB struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:100;not null"`
	Age       int            `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UsersDB) TableName() string {
	return "users"
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	db.AutoMigrate(&ItemDB{}, &UsersDB{})
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Save(item *item.Item) error {
	dbItem := ItemDB{
		Name:        item.Name,
		Price:       item.Price,
		Description: item.Description,
		Stock:       item.Stock,
	}

	err := r.db.Create(&dbItem).Error
	if err != nil {
		return err
	}
	item.ID = dbItem.ID

	return nil
}

func (r *ItemRepository) GetItemByID(id int) (item.Item, error) {
	itemModel := &ItemDB{}
	//select * from items where deleted_at is null and id = id;
	err := r.db.First(itemModel, id).Error
	if err != nil {
		return item.Item{}, err
	}

	return item.Item{
		ID:          itemModel.ID,
		Name:        itemModel.Name,
		Price:       itemModel.Price,
		Description: itemModel.Description,
		Stock:       itemModel.Stock,
	}, nil
}

func (r *ItemRepository) GetItems() ([]item.Item, error) {
	items := []ItemDB{}
	//select * from items where deleted_at is null;
	err := r.db.Find(&items).Order("id").Error

	if err != nil {
		return nil, err
	}

	var itemsList []item.Item
	for _, itemModel := range items {
		itemsList = append(itemsList, item.Item{
			ID:          itemModel.ID,
			Name:        itemModel.Name,
			Price:       itemModel.Price,
			Description: itemModel.Description,
			Stock:       itemModel.Stock,
		})
	}

	return itemsList, nil
}

func (r *ItemRepository) DeleteItem(id int) error {
	err := r.db.Delete(&ItemDB{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
