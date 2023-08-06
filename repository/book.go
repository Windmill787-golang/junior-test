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

	if err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Genre, &book.Author, &book.PageCount, &book.Year, &book.Price, &book.CreatedAt, &book.UpdatedAt, &book.UserId); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) GetBooks() ([]*entities.Book, error) {
	rows, err := r.db.Query(fmt.Sprintf("SELECT * FROM %s", booksTable))
	if err != nil {
		return nil, err
	}

	books := make([]*entities.Book, 0)
	for rows.Next() {
		book := entities.NewBook()
		if err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Genre, &book.Author, &book.PageCount, &book.Year, &book.Price, &book.CreatedAt, &book.UpdatedAt, &book.UserId); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) CreateBook(book entities.Book) (int, error) {
	sql := fmt.Sprintf("INSERT INTO %s "+
		"(title, description, genre, author, page_count, year, price, user_id) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) "+
		"RETURNING id",
		booksTable,
	)

	var id int
	row := r.db.QueryRow(sql, book.Title, book.Description, book.Genre, book.Author, book.PageCount, book.Year, book.Price, book.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BookRepository) UpdateBook(book entities.Book) error {
	sql := fmt.Sprintf("UPDATE %s SET "+
		"title=$1, description=$2, genre=$3, author=$4, page_count=$5,"+
		"year=$6, price=$7, updated_at=now() WHERE id=$8",
		booksTable,
	)

	_, err := r.db.Exec(sql, book.Title, book.Description, book.Genre, book.Author, book.PageCount, book.Year, book.Price, book.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) DeleteBook(id int) error {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id=$1", booksTable)
	_, err := r.db.Exec(sql, id)
	if err != nil {
		return err
	}

	return nil
}
