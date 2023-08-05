package entities

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Author      string `json:"author"`
	PageCount   int    `json:"page_count"`
	ReleaseDate string `json:"release_date"`
	Price       int    `json:"price"`
}

func NewBook() *Book {
	return &Book{}
}
