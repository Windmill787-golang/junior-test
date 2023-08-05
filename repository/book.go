package repository

import (
	"database/sql"

	"github.com/Windmill787-golang/junior-test/entities"
)

type BookRepository struct {
	db *sql.DB
}

func NewBook(db *sql.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) GetBook(id int) (*entities.Book, error) {
	return &entities.Book{}, nil
}
