package database

import (
	"rest_api_test/models"
)

// GetAllAuthor from database
func GetAllAuthor() []models.Author {
	var authors []models.Author

	rows, err := getDB().Query("SELECT id,name FROM authors where status=1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var author models.Author
		rows.Scan(&author.ID, &author.Name)
		authors = append(authors, author)
	}

	defer getDB().Close()

	return authors
}

// GetAuthorByID from database
func GetAuthorByID(id string) models.Author {
	var author models.Author

	rows, err := getDB().Query("SELECT id,name FROM authors where status=1 and id=" + id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		rows.Scan(&author.ID, &author.Name)
	}

	defer getDB().Close()

	return author

}

// SaveAuthor into database
func SaveAuthor(author models.Author) models.Author {
	insert, err := getDB().Exec("INSERT INTO authors (name) VALUES ('" + author.Name + "')")
	if err != nil {
		panic(err.Error())
	}

	lastID, err := insert.LastInsertId()

	defer getDB().Close()

	return models.Author{
		ID:   lastID,
		Name: author.Name,
	}
}

// DeleteAuthorByID if ID exists
func DeleteAuthorByID(id string) bool {
	res, err := getDB().Exec("update authors set status=0 where id=" + id)
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

// UpdateAuthorByID in database
func UpdateAuthorByID(author models.Author, id string) bool {
	res, err := getDB().Exec("update authors set name='" + author.Name + "' where id=" + id)
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
