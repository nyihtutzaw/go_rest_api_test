package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api_test/authservice"
	"rest_api_test/database"
	"rest_api_test/models"
	"rest_api_test/utils"
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

	r.ParseMultipartForm(10 << 20)

	err, fileName := utils.UploadFile(r, "image", "books")

	if err != nil {
		fmt.Printf(err.Error())
	}

	authorID, err := strconv.ParseInt(r.FormValue("authorID"), 10, 64)

	book := database.SaveBook(models.Book{
		Name:     r.FormValue("name"),
		AuthorID: authorID,
		Image:    fileName,
	})

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

	book := database.GetBookByID(params["id"])

	utils.RemoveFile("books", book.Image)

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
