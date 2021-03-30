package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api_test/authservice"
	"rest_api_test/database"
	"rest_api_test/models"
	"rest_api_test/utils"

	"github.com/gorilla/mux"
)

func getAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	authors = database.GetAllAuthor()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&models.AuthorsResponseType{
		Data:    authors,
		Message: "All Authors",
	})
}

func getAuthor(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	author := database.GetAuthorByID(params["id"])
	if author.ID == 0 {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(&models.AuthorResponseType{
			Data:    author,
			Message: "Author by ID " + params["id"],
		})
	}

}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseMultipartForm(10 << 20)

	err, fileName := utils.UploadFile(r, "image", "authors")

	if err != nil {
		fmt.Printf(err.Error())
	}

	author := database.SaveAuthor(models.Author{
		Name:  r.FormValue("name"),
		Image: fileName,
	})

	json.NewEncoder(w).Encode(&models.AuthorResponseType{
		Data:    author,
		Message: "New Author is saved",
	})
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var author models.Author
	_ = json.NewDecoder(r.Body).Decode(&author)

	result := database.UpdateAuthorByID(author, params["id"])
	if !result {
		w.WriteHeader(404)
	} else {
		author = database.GetAuthorByID(params["id"])

		json.NewEncoder(w).Encode(&models.AuthorResponseType{
			Data:    author,
			Message: "Author is updated",
		})
	}

}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	author := database.GetAuthorByID(params["id"])

	utils.RemoveFile("authors", author.Image)

	result := database.DeleteAuthorByID(params["id"])
	if !result {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(&models.ResponseType{
			Message: "Author by ID " + params["id"] + " is deleted",
		})
	}

}

func authorHandler(r *mux.Router) {

	e := r.PathPrefix("/authors").Subrouter()
	e.Use(authservice.Middleware)
	e.HandleFunc("", getAuthors).Methods("GET")

	s := r.PathPrefix("/author").Subrouter()
	s.Use(authservice.Middleware)
	s.HandleFunc("/{id}", getAuthor).Methods("GET")
	s.HandleFunc("/store", createAuthor).Methods("POST")
	s.HandleFunc("/{id}", updateAuthor).Methods("PUT")
	s.HandleFunc("/{id}", deleteAuthor).Methods("DELETE")

}
