package database

import (
	"fmt"
	"rest_api_test/models"
)

// RegisterUser from database
func RegisterUser(userData models.User) models.User {

	result := GetDB().Create(&userData)
	if result.Error != nil {
		fmt.Printf(result.Error.Error())
	}

	return userData

}

// CheckUserExist from database
func CheckUserExist(username string) int {

	result := GetDB().Where("username = ?", username).Find(&models.User{})

	return int(result.RowsAffected)
}

// GetUserByUsername from database
func GetUserByUsername(username string) models.User {
	var user models.User

	GetDB().Where("username = ?", username).Find(&user)
	return user

}

// GetUserByIDAnndUsername from database
func GetUserByIDAnndUsername(userID float64, username string) models.User {
	var user models.User

	GetDB().Where("username = ?", username).Where("id=?", userID).Find(&user)
	return user

}
