package database

import (
	"database/sql"
	"rest_api_test/models"
	"strconv"
)

type resultData struct {
	authorName string
}

// GetAllBook from database
func GetAllBook() []models.Book {
	var books []models.Book

	rows, err := getDB().Query("SELECT books.id as ID ,books.name,books.authorID,authors.name as authorName FROM books,authors where books.authorID=authors.id and books.status=1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var book models.Book
		book = getBookObj(rows)
		books = append(books, book)
	}

	defer getDB().Close()

	return books
}

// GetAllBookByAuthorID from database
func GetAllBookByAuthorID(authorID string) []models.Book {
	var books []models.Book

	rows, err := getDB().Query("SELECT books.id as ID ,books.name,books.authorID,authors.name as authorName FROM books,authors where books.authorID=authors.id and books.status=1 and books.authorID=" + authorID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var book models.Book
		book = getBookObj(rows)
		books = append(books, book)
	}

	defer getDB().Close()

	return books
}

// GetBookByID from database
func GetBookByID(id string) models.Book {

	var book models.Book

	rows, err := getDB().Query("SELECT books.id as ID ,books.name,books.authorID,authors.name as authorName FROM books,authors where books.authorID=authors.id and books.status=1 and books.id=" + id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		book = getBookObj(rows)
	}

	defer getDB().Close()

	return book

}

// SaveBook into database
func SaveBook(book models.Book) models.Book {
	insert, err := getDB().Exec("INSERT INTO books (name,authorID) VALUES ('" + book.Name + "'," + strconv.Itoa(int(book.AuthorID)) + ")")
	if err != nil {
		panic(err.Error())
	}

	lastID, err := insert.LastInsertId()

	defer getDB().Close()

	return models.Book{
		ID:       lastID,
		Name:     book.Name,
		AuthorID: book.AuthorID,
	}
}

// DeleteBookByID if exists
func DeleteBookByID(id string) bool {
	res, err := getDB().Exec("update books set status=0 where id=" + id)
	if err != nil {
		panic(err.Error())
	}
	row, err := res.RowsAffected()
	defer getDB().Close()

	if row > 0 {
		return true
	}

	return false
}

// UpdateBookByID if exists
func UpdateBookByID(book models.Book, id string) bool {
	res, err := getDB().Exec("update books set name='" + book.Name + "',authorID=" + strconv.Itoa(int(book.AuthorID)) + " where id=" + id)
	if err != nil {
		panic(err.Error())
	}
	row, err := res.RowsAffected()
	defer getDB().Close()

	if row > 0 {
		return true
	}

	return false
}

func getBookObj(rows *sql.Rows) models.Book {
	var customData resultData
	var book models.Book

	rows.Scan(&book.ID, &book.Name, &book.AuthorID, &customData.authorName)

	book.Author = models.Author{
		ID:   book.AuthorID,
		Name: customData.authorName,
	}
	return book

}
