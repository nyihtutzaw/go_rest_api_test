package database

import (
	"fmt"
	"rest_api_test/config"
	"rest_api_test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDB
func GetDB() *gorm.DB {

	dsn := config.GETEnvVariable("DB_USERNAME") + ":" + config.GETEnvVariable("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/" + config.GETEnvVariable("DB_NAME") + "?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Printf(err.Error())
	}

	return db

}

func MigrateTables() {

	err := GetDB().AutoMigrate(&models.User{}, &models.Author{}, &models.Book{})
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Migration success")
	}
}
