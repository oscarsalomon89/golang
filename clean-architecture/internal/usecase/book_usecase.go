package usecase

import (
	"errors"

	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
}

func NewBookUsecase(bookRepo domain.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: bookRepo,
	}
}

func (u *bookUsecase) GetBooks() []domain.Book {
	return u.bookRepository.GetBooks()
}

func (u *bookUsecase) GetBook(id int) *domain.Book {
	return u.bookRepository.GetBook(id)
}

func (u *bookUsecase) AddBook(book domain.Book) (*domain.Book, error) {
	books := u.bookRepository.GetBooks()
	for _, b := range books {
		if b.ID == book.ID {
			return nil, errors.New("book already exist")
		}
	}

	result := u.bookRepository.AddBook(book)

	return result, nil
}

func (u *bookUsecase) UpdateBook(id int, book domain.Book) (*domain.Book, error) {
	books := u.bookRepository.GetBooks()
	for key, b := range books {
		if b.ID == id {
			result := u.bookRepository.UpdateBook(key, book)
			return result, nil
		}
	}

	return nil, errors.New("book not found")
}

func (u *bookUsecase) DeleteBook(id int) error {
	return u.DeleteBook(id)
}
