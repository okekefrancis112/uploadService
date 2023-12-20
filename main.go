// main.go
package main

import (
	// "os"
	"uploadService/mod/routes"
	"uploadService/mod/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	database.InitDatabase()
	defer database.DBConn.Close()
	routes.SetupRoutes(app)
	app.Listen(":3000")
}

