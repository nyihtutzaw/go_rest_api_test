package models

// Book models
type Book struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	AuthorID int64  `json:"authorID"`
	Author   Author `gorm:"AuthorID" json:"author"`
	Image    string `json:"image"`
}

// BooksResponseType Model
type BooksResponseType struct {
	Message string `json:"message"`
	Data    []Book `json:"data"`
}

// BookResponseType Model
type BookResponseType struct {
	Message string `json:"message"`
	Data    Book   `json:"data"`
}
