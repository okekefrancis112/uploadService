package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"os"
	"uploadService/mod/models"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	// Use the database URL from the environment variable
	// database.DBConn, err = gorm.Open("sqlite3", os.Getenv("DB_URL"))
	DBConn, err = gorm.Open("sqlite3", "videos.db")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}

	// Auto Migrate - Creates tables based on struct models
	DBConn.AutoMigrate(&models.Video{})

	fmt.Println("Connected to the database successfully.")
}