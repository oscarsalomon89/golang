package repository

import "github.com/teamcubation/neocamp-meli/rest/restful/entity"

var db []entity.Book
var id int = 1

type BookRepository struct {
}

func (r *BookRepository) GetBooks() []entity.Book {
	return db
}

func (r *BookRepository) GetBook(id int) entity.Book {
	for _, v := range db {
		if v.ID == id {
			return v
		}
	}
	return entity.Book{}
}

func (r *BookRepository) GetBookByName(name string) entity.Book {
	for _, v := range db {
		if v.Title == name {
			return v
		}
	}
	return entity.Book{}
}

func (r *BookRepository) AddBook(book entity.Book) entity.Book {
	book.ID = id
	db = append(db, book)

	id++

	return book
}

func (r *BookRepository) UpdateBook(book entity.Book) entity.Book {
	for i, v := range db {
		if v.ID == book.ID {
			db[i] = book
		}
	}

	return book
}

func (r *BookRepository) DeleteBook(id int) entity.Book {
	for i, v := range db {
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	return entity.Book{}
}
