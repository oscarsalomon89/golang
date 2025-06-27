package item

type Repository interface {
	Save(item *Item) error
	GetItemByID(id int) (Item, error)
	GetItems() ([]Item, error)
	DeleteItem(id int) error
}
