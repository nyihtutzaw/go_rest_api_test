package routes

import (
	"encoding/json"
	"net/http"
	"rest_api_test/authservice"
	"rest_api_test/database"
	"rest_api_test/models"
	"strconv"

	"github.com/gorilla/mux"
)

func getBooks(w http.ResponseWriter, r *http.Request) {

	var books []models.Book

	authorID, checkAuthorID := strconv.Atoi(r.URL.Query().Get("authorID"))

	if checkAuthorID == nil {
		books = database.GetAllBookByAuthorID(strconv.Itoa(int(authorID)))
	} else {
		books = database.GetAllBook()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&models.BooksResponseType{
		Data:    books,
		Message: "All Books",
	})
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	book := database.GetBookByID(params["id"])
	if book.ID == 0 {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(&models.BookResponseType{
			Data:    book,
			Message: "Book by ID " + params["id"],
		})
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book = database.SaveBook(book)

	book = database.GetBookByID(strconv.Itoa(int(book.ID)))

	json.NewEncoder(w).Encode(&models.BookResponseType{
		Data:    book,
		Message: "New Book is saved",
	})
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	result := database.UpdateBookByID(book, params["id"])
	if !result {
		w.WriteHeader(404)
	} else {
		book = database.GetBookByID(params["id"])

		json.NewEncoder(w).Encode(&models.BookResponseType{
			Data:    book,
			Message: "Book is updated",
		})
	}

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result := database.DeleteBookByID(params["id"])
	if !result {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(&models.ResponseType{
			Message: "Book by ID " + params["id"] + " is deleted",
		})
	}

}

func bookHandler(r *mux.Router) {
	e := r.PathPrefix("/books").Subrouter()
	e.Use(authservice.Middleware)
	e.HandleFunc("", getBooks).Methods("GET")

	s := r.PathPrefix("/book").Subrouter()
	s.Use(authservice.Middleware)
	s.HandleFunc("/{id}", getBook).Methods("GET")
	s.HandleFunc("/store", createBook).Methods("POST")
	s.HandleFunc("/{id}", updateBook).Methods("PUT")
	s.HandleFunc("/{id}", deleteBook).Methods("DELETE")

}
