package entities

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Author      string `json:"author" binding:"required"`
	PageCount   int    `json:"page_count" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UserId      int    `json:"user_id"`
}

func NewBook() *Book {
	return &Book{}
}
