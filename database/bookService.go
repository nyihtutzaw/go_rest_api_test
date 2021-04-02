package database

import (
	"fmt"
	"rest_api_test/models"
	"rest_api_test/utils"
)

type resultData struct {
	authorName  string
	authorImage string
}

// GetAllBook from database
func GetAllBook() []models.Book {
	books := []models.Book{}
	GetDB().Preload("Author").Find(&books)
	for i, v := range books {
		books[i].Image = utils.GetFullName(v.Image, "books")
		books[i].Author.Image = utils.GetFullName(v.Author.Image, "authors")
	}

	return books
}

// GetAllBookByAuthorID from database
func GetAllBookByAuthorID(authorID string) []models.Book {

	books := []models.Book{}
	GetDB().Preload("Author").Where("author_id=?", authorID).Find(&books)
	for i, v := range books {
		books[i].Image = utils.GetFullName(v.Image, "books")
		books[i].Author.Image = utils.GetFullName(v.Author.Image, "authors")
	}
	return books
}

// GetBookByID from database
func GetBookByID(id string) models.Book {

	var book models.Book
	GetDB().Preload("Author").Where("id=?", id).First(&book)
	book.Image = utils.GetFullName(book.Image, "books")
	book.Author.Image = utils.GetFullName(book.Author.Image, "authors")
	return book

}

// SaveBook into database
func SaveBook(book models.Book) models.Book {
	result := GetDB().Create(&book)
	if result.Error != nil {
		fmt.Printf(result.Error.Error())
	}
	return book
}

// DeleteBookByID if exists
func DeleteBookByID(id string) bool {
	result := GetDB().Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// UpdateBookByID if exists
func UpdateBookByID(book models.Book, id string) bool {
	results := GetDB().Where("id=?", id).UpdateColumns(book)
	if results.RowsAffected > 0 {
		return true
	}
	return false

}
