package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Windmill787-golang/junior-test/internal/entities"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) CreateUser(user entities.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id",
		usersTable,
	)

	var id int
	row := r.db.QueryRow(query, user.Username, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepository) GetUserId(username, passwordHash string) (int, error) {
	row := r.db.QueryRow(fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", usersTable), username, passwordHash)

	var id int

	if err := row.Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return id, nil
}
