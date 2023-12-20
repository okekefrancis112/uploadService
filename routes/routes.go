package routes

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/gofiber/fiber/v2"
	"uploadService/mod/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/v1/file", controllers.FileUpload)
	app.Post("/api/v1/remote", controllers.RemoteUpload)
	app.Get("/api/v1/files", controllers.GetFiles)
	app.Get("/api/v1/file/:id", controllers.GetFile)
	app.Delete("/api/v1/file/:id", controllers.DeleteFile)
}