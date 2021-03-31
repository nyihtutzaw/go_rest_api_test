package database

import (
	"fmt"
	"rest_api_test/models"
	"rest_api_test/utils"
)

// GetAllAuthor from database
func GetAllAuthor() []models.Author {

	authors := []models.Author{}
	GetDB().Find(&authors)
	for i, v := range authors {
		authors[i].Image = utils.GetFullName(v.Image, "authors")
	}
	return authors
}

// GetAuthorByID from database
func GetAuthorByID(id string) models.Author {
	var author models.Author
	GetDB().Where("id=?", id).First(&author)
	author.Image = utils.GetFullName(author.Image, "authors")
	return author
}

// SaveAuthor into database
func SaveAuthor(author models.Author) models.Author {
	result := GetDB().Create(&author)
	if result.Error != nil {
		fmt.Printf(result.Error.Error())
	}
	return author
}

// DeleteAuthorByID if ID exists
func DeleteAuthorByID(id string) bool {
	result := GetDB().Delete(&models.Author{}, id)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// UpdateAuthorByID in database
func UpdateAuthorByID(author models.Author, id string) bool {
	results := GetDB().Where("id=?", id).UpdateColumns(author)
	if results.RowsAffected > 0 {
		return true
	}
	return false
}
