package authservice

import (
	"fmt"
	"net/http"
	"rest_api_test/database"
	"rest_api_test/models"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//CreateToken after user login
func CreateToken(user models.User) string {
	var err error
	//Creating Access Token

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["username"] = user.Username
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("alliswell"))
	if err != nil {
		panic(err.Error())
	}
	return token
}

//ExtractToken from barrier
func ExtractToken(r *http.Request) (string, bool) {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1], true
	}

	return "", false
}

//CheckToken claims in database
func CheckToken(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("alliswell"), nil
	})
	if err != nil {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user := database.GetUserByIDAnndUsername(claims["user_id"].(float64), claims["username"].(string))
		if user.ID > 0 {
			return true
		}
	}
	return false
}
