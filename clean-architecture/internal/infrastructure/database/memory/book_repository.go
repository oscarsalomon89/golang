package memory

import (
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain/model"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain/repository"
)

type bookRepository struct {
	db []model.Book
}

func NewBookRepository(db []model.Book) repository.BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) GetAllBooks() []model.Book {
	return r.db
}

func (r *bookRepository) SaveBook(book model.Book) model.Book {
	r.db = append(r.db, book)

	return book
}
