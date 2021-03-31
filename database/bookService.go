package database

import (
	"database/sql"
	"rest_api_test/models"
	"rest_api_test/utils"
)

type resultData struct {
	authorName  string
	authorImage string
}

// GetAllBook from database
func GetAllBook() []models.Book {
	var books []models.Book

	// rows, err := getDB().Query("SELECT books.id as ID ,books.name,books.authorID,authors.name as authorName,authors.image as authorImage,books.image FROM books,authors where books.authorID=authors.id and books.status=1")
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// for rows.Next() {
	// 	var book models.Book
	// 	book = getBookObj(rows)

	// 	books = append(books, book)
	// }

	return books
}

// GetAllBookByAuthorID from database
func GetAllBookByAuthorID(authorID string) []models.Book {
	var books []models.Book

	// rows, err := getDB().Query("SELECT books.id as ID ,books.name,books.authorID,authors.name as authorName,authors.image as authorImage,books.image FROM books,authors where books.authorID=authors.id and books.status=1 and books.authorID=" + authorID)
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// for rows.Next() {
	// 	var book models.Book
	// 	book = getBookObj(rows)

	// 	books = append(books, book)
	// }

	return books
}

// GetBookByID from database
func GetBookByID(id string) models.Book {

	var book models.Book

	// rows, err := getDB().Query("SELECT books.id as ID,books.name,books.authorID,authors.name as authorName,authors.image as authorImage,books.image FROM books,authors where books.authorID=authors.id and books.status=1 and books.id=" + id)
	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// for rows.Next() {
	// 	book = getBookObj(rows)

	// }

	return book

}

// SaveBook into database
func SaveBook(book models.Book) models.Book {
	// insert, err := getDB().Exec("INSERT INTO books (name,authorID,image) VALUES ('" + book.Name + "'," + strconv.Itoa(int(book.AuthorID)) + ",'" + book.Image + "')")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// lastID, err := insert.LastInsertId()

	// return models.Book{
	// 	ID:       lastID,
	// 	Name:     book.Name,
	// 	AuthorID: book.AuthorID,
	// }

	return models.Book{}
}

// DeleteBookByID if exists
func DeleteBookByID(id string) bool {
	// res, err := getDB().Exec("update books set status=0 where id=" + id)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// row, err := res.RowsAffected()

	// if row > 0 {
	// 	return true
	// }

	return false
}

// UpdateBookByID if exists
func UpdateBookByID(book models.Book, id string) bool {
	// res, err := getDB().Exec("update books set name='" + book.Name + "',authorID=" + strconv.Itoa(int(book.AuthorID)) + " where id=" + id)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// row, err := res.RowsAffected()

	// if row > 0 {
	// 	return true
	// }

	return false
}

func getBookObj(rows *sql.Rows) models.Book {
	var customData resultData
	var book models.Book

	rows.Scan(&book.ID, &book.Name, &book.AuthorID, &customData.authorName, &customData.authorImage, &book.Image)

	customData.authorImage = utils.GetFullName(customData.authorImage, "authors")

	book.Author = models.Author{
		ID:    book.AuthorID,
		Name:  customData.authorName,
		Image: customData.authorImage,
	}

	book.Image = utils.GetFullName(book.Image, "books")
	return book
}
