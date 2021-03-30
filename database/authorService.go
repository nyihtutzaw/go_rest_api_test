package database

import (
	"rest_api_test/models"
	"rest_api_test/utils"
)

// GetAllAuthor from database
func GetAllAuthor() []models.Author {

	authors := []models.Author{}

	rows, err := getDB().Query("SELECT id,name,image FROM authors where status=1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		var author models.Author
		rows.Scan(&author.ID, &author.Name, &author.Image)
		author.Image = utils.GetFullName(author.Image, "authors")
		authors = append(authors, author)
	}

	defer getDB().Close()

	return authors
}

// GetAuthorByID from database
func GetAuthorByID(id string) models.Author {
	var author models.Author

	rows, err := getDB().Query("SELECT id,name,image FROM authors where status=1 and id=" + id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		rows.Scan(&author.ID, &author.Name, &author.Image)
		author.Image = utils.GetFullName(author.Image, "authors")
	}

	defer getDB().Close()

	return author

}

// SaveAuthor into database
func SaveAuthor(author models.Author) models.Author {
	insert, err := getDB().Exec("INSERT INTO authors (name,image) VALUES ('" + author.Name + "','" + author.Image + "')")
	if err != nil {
		panic(err.Error())
	}

	lastID, err := insert.LastInsertId()

	defer getDB().Close()

	return models.Author{
		ID:    lastID,
		Name:  author.Name,
		Image: author.Image,
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
