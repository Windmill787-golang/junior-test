package repository

import (
	"database/sql"

	"github.com/Windmill787-golang/junior-test/internal/entities"
)

type Book interface {
	GetBook(id int) (*entities.Book, error)
	GetBooks() ([]*entities.Book, error)
	CreateBook(book entities.Book) (int, error)
	UpdateBook(book entities.Book) error
	DeleteBook(id int) error
}

type Auth interface {
	CreateUser(user entities.User) (int, error)
	GetUserId(username, passwordHash string) (int, error)
}

type Repository struct {
	Book
	Auth
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book: NewBookRepository(db),
		Auth: NewAuthRepository(db),
	}
}
