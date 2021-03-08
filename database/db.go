package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() *sql.DB {
	connection, err := sql.Open("mysql", "nyihtutzaw:00270027nhz@/go_rest_test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	return connection
}
