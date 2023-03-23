package repository

import (
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"
)

var db []domain.Book

type bookRepository struct{}

func NewBookRepository() domain.BookRepository {
	db = []domain.Book{
		{
			ID:     1,
			Title:  "Dune",
			Price:  1965,
			Author: "Frank Herbert",
		},
		{
			ID:     2,
			Title:  "Cita con Rama",
			Price:  1974,
			Author: "Arthur C. Clarke",
		},
		{
			ID:     3,
			Title:  "Un guijarro en el cielo",
			Price:  500,
			Author: "Isaac Asimov",
		},
	}

	return &bookRepository{}
}

func (u *bookRepository) GetBooks() []domain.Book {
	return db
}

func (u *bookRepository) GetBook(id int) *domain.Book {
	for _, book := range db {
		if book.ID == id {
			return &book
		}
	}

	return nil
}

func (u *bookRepository) AddBook(book domain.Book) *domain.Book {
	db = append(db, book)

	return &book
}

func (u *bookRepository) UpdateBook(key int, book domain.Book) *domain.Book {
	db[key] = book

	return &book
}

func (u *bookRepository) DeleteBook(id int) error {
	for i, v := range db {
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	return nil
}
