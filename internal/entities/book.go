package entities

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	Author      string `json:"author" validate:"required"`
	PageCount   int    `json:"page_count" validate:"required"`
	Year        int    `json:"year" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UserId      int    `json:"user_id"` //TODO: validate
}

func NewBook() *Book {
	return &Book{}
}
