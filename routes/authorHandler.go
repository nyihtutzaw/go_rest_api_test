package routes

import (
	"encoding/json"
	"net/http"
	"rest_api_test/database"
	"rest_api_test/models"

	"github.com/gorilla/mux"
)

var authors []models.Author

func getAuthors(w http.ResponseWriter, r *http.Request) {

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
	var author models.Author
	_ = json.NewDecoder(r.Body).Decode(&author)

	author = database.SaveAuthor(author)

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

	r.HandleFunc("/authors", getAuthors).Methods("GET")

	s := r.PathPrefix("/author").Subrouter()

	s.HandleFunc("/{id}", getAuthor).Methods("GET")
	s.HandleFunc("/store", createAuthor).Methods("POST")
	s.HandleFunc("/{id}", updateAuthor).Methods("PUT")
	s.HandleFunc("/{id}", deleteAuthor).Methods("DELETE")

}
