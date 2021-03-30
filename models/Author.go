package models

// Author models
type Author struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// AuthorsResponseType Model
type AuthorsResponseType struct {
	Message string   `json:"message"`
	Data    []Author `json:"data"`
}

// AuthorResponseType Model
type AuthorResponseType struct {
	Message string `json:"message"`
	Data    Author `json:"data"`
}
