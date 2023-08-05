package service

import (
	"github.com/Windmill787-golang/junior-test/entities"
	"github.com/Windmill787-golang/junior-test/repository"
)

type Book interface {
	GetBook(id int) (*entities.Book, error)
}

type Service struct {
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repo.Book),
	}
}
