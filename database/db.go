package database

import (
	"database/sql"
	"rest_api_test/config"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() *sql.DB {
	connection, err := sql.Open(config.GETEnvVariable("DB"), config.GETEnvVariable("DB_USERNAME")+":"+config.GETEnvVariable("DB_PASSWORD")+"@/"+config.GETEnvVariable("DB_NAME"))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	return connection
}
