package usecase

import (
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain/model"
	"github.com/teamcubation/neocamp-meli/clean-architecture/internal/domain/repository"
)

type BookUsecase interface {
	GetAllBooks() []model.Book
	SaveBook(book model.Book) model.Book
}

type bookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(repo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		repo: repo,
	}
}

func (u *bookUsecase) GetAllBooks() []model.Book {
	return u.repo.GetAllBooks()
}

func (u *bookUsecase) SaveBook(book model.Book) model.Book {
	return u.repo.SaveBook(book)
}
