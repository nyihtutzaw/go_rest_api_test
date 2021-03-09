package database

import (
	"rest_api_test/models"
	"strconv"
)

// RegisterUser from database
func RegisterUser(userData models.UserBody) models.User {

	insert, err := getDB().Exec("INSERT INTO users (username,password) VALUES ('" + userData.UserData.Username + "','" + userData.Password + "')")
	if err != nil {
		panic(err.Error())
	}

	lastID, err := insert.LastInsertId()

	defer getDB().Close()

	return models.User{
		ID:       lastID,
		Username: userData.UserData.Username,
	}

}

// CheckUserExist from database
func CheckUserExist(username string) int {
	rows, err := getDB().Query("SELECT * FROM users where username='" + username + "'")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	count := 0
	for rows.Next() {
		count++
	}

	return count
}

// GetUserByUsername from database
func GetUserByUsername(username string) models.UserBody {
	var user models.User
	var password string

	rows, err := getDB().Query("SELECT * FROM users where username='" + username + "'")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &password)
	}

	defer getDB().Close()

	return models.UserBody{
		UserData: user,
		Password: password,
	}

}

// GetUserByIDAnndUsername from database
func GetUserByIDAnndUsername(userID float64, username string) models.User {
	var user models.User

	id := strconv.Itoa(int(userID))

	rows, err := getDB().Query("SELECT id,username FROM users where id=" + id + " and username='" + username + "'")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for rows.Next() {
		rows.Scan(&user.ID, &user.Username)
	}

	defer getDB().Close()

	return user

}
