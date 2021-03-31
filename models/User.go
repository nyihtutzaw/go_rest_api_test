package models

// User models
type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// // UserBody models is user model that inclues password
// type UserBody struct {
// 	UserData User
// 	Password string
// }

// UserResponseType models
type UserResponseType struct {
	User    User   `json:"user"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
