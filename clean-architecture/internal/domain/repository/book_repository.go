package repository

import "github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain/model"

type BookRepository interface {
	GetAllBooks() []model.Book
	SaveBook(book model.Book) model.Book
}
