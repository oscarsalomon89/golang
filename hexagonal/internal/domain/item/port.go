package item

type Repository interface {
	Save(item Item) error
}
