package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/teamcubation/neocamp-meli/rest/restful/entity"
)

type BookModel struct {
	ID        int       `db:"id"`
	Author    string    `db:"author"`
	Title     string    `db:"title"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetBooks(name string) []entity.Book {
	books := []BookModel{}
	if name == "" {
		err := r.db.Select(&books, "SELECT * FROM books")
		if err != nil {
			return []entity.Book{}
		}
	} else {
		err := r.db.Select(&books, "SELECT * FROM books WHERE title = ?", name)
		if err != nil {
			return []entity.Book{}
		}
	}

	booksEntity := []entity.Book{}
	for _, book := range books {
		booksEntity = append(booksEntity, entity.Book{
			ID:     book.ID,
			Author: book.Author,
			Title:  book.Title,
			Price:  book.Price,
		})
	}

	return booksEntity
}

func (r *BookRepository) GetBook(id int) entity.Book {
	book := BookModel{}
	err := r.db.Get(&book, "SELECT id, author, title, price FROM books WHERE id = ?", id)
	if err != nil {
		return entity.Book{}
	}
	return entity.Book{
		ID:     book.ID,
		Author: book.Author,
		Title:  book.Title,
		Price:  book.Price,
	}
}

func (r *BookRepository) GetBookByName(name string) entity.Book {
	book := entity.Book{}
	r.db.Get(&book, "SELECT id, author, title, price FROM books WHERE title = ?", name)
	return book
}

func (r *BookRepository) AddBook(book entity.Book) (entity.Book, error) {
	result, err := r.db.Exec("INSERT INTO books (author, title, price) VALUES (?, ?, ?)", book.Author, book.Title, book.Price)
	if err != nil {
		return entity.Book{}, err
	}

	id, _ := result.LastInsertId()
	book.ID = int(id)
	return book, nil
}

func (r *BookRepository) UpdateBook(book entity.Book) entity.Book {
	_, err := r.db.Exec("UPDATE books SET author = ?, title = ?, price = ?, updated_at = NOW() WHERE id = ?", book.Author, book.Title, book.Price, book.ID)
	if err != nil {
		return entity.Book{}
	}
	return book
}

func (r *BookRepository) DeleteBook(id int) entity.Book {
	book := entity.Book{}
	r.db.Exec("DELETE FROM books WHERE id = ?", id)
	return book
}

// func (r *BookRepository) GetBooks() []entity.Book {
// 	var books []entity.Book
// 	r.db.Find(&books)
// 	return books
// }

// func (r *BookRepository) GetBook(id int) entity.Book {
// 	var book entity.Book
// 	r.db.First(&book, id)
// 	return book
// }

// func (r *BookRepository) GetBookByName(name string) entity.Book {
// 	var book entity.Book
// 	r.db.Where("title = ?", name).First(&book)
// 	return book
// }

// func (r *BookRepository) AddBook(book entity.Book) (entity.Book, error) {
// 	if err := r.db.Create(&book).Error; err != nil {
// 		return entity.Book{}, err
// 	}
// 	return book, nil
// }

// func (r *BookRepository) UpdateBook(book entity.Book) entity.Book {
// 	r.db.Save(&book)
// 	return book
// }

// func (r *BookRepository) DeleteBook(id int) entity.Book {
// 	var book entity.Book
// 	r.db.Delete(&book, id)
// 	return book
// }
