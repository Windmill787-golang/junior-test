package repository

import (
	"database/sql"

	"github.com/Windmill787-golang/junior-test/entities"
)

type Book interface {
	GetBook(id int) (*entities.Book, error)
	GetBooks() ([]*entities.Book, error)
	CreateBook(entities.Book) (int, error)
}

type Repository struct {
	Book
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book: NewBookRepository(db),
	}
}
