package routes

import (
	"encoding/json"
	"net/http"
	"rest_api_test/authservice"
	"rest_api_test/database"
	"rest_api_test/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type authData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t authData
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(t.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	if database.CheckUserExist(t.Username) > 0 {
		json.NewEncoder(w).Encode(models.ResponseType{
			Message: "Username already existed",
		})
	} else {
		user := database.RegisterUser(models.User{
			Username: t.Username,
			Password: string(hashedPassword),
		})
		json.NewEncoder(w).Encode(models.User{
			Username: user.Username,
			ID:       user.ID,
		})
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t authData
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	user := database.GetUserByUsername(t.Username)

	if user.ID > 0 {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(t.Password))
		if err == nil {
			json.NewEncoder(w).Encode(models.UserResponseType{
				Message: "Login success",
				User: models.User{
					Username: user.Username,
					ID:       user.ID,
				},
				Token: authservice.CreateToken(user),
			})
		} else {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(models.ResponseType{
				Message: "Password not match",
			})
		}

	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(models.ResponseType{
			Message: "Username not found",
		})

	}

}

func userHandler(r *mux.Router) {

	s := r.PathPrefix("/user").Subrouter()

	s.HandleFunc("/register", register).Methods("POST")
	s.HandleFunc("/login", login).Methods("POST")

}
