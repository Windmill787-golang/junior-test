package repository

import (
	"database/sql"
	"fmt"

	"github.com/Windmill787-golang/junior-test/entities"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) GetBook(id int) (*entities.Book, error) {
	row := r.db.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id = $1", booksTable), id)

	book := entities.NewBook()

	if err := row.Scan(book.ID, book.Title, book.Description, book.Genre,
		book.Genre, book.Author, book.PageCount, book.ReleaseDate,
		book.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return book, nil
}
