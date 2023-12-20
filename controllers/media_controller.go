package controllers

import (
	"uploadService/mod/handlers"
    "uploadService/mod/database"
	"uploadService/mod/models"
	"uploadService/mod/services"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func FileUpload(c *fiber.Ctx) error {
    db := database.DBConn
	formHeader, err := c.FormFile("file")
	if err != nil {
		return handlers.HandleError(c, http.StatusInternalServerError, "Select a file to upload", err)
	}

	formFile, err := formHeader.Open()
	if err != nil {
		return handlers.HandleError(c, http.StatusInternalServerError, "Error opening file", err)
	}

	uploadUrl, err := services.NewMediaUpload().FileUpload(models.File{File: formFile})
	if err != nil {
		return handlers.HandleError(c, http.StatusInternalServerError, "Error uploading file", err)
	}

	// Insert video information into the database
	video := models.Video{
		Video_Url:  uploadUrl,
	}

	db.Create(&video)

	// Return the uploaded video information as JSON response
	return c.JSON(video)
}

func RemoteUpload(c *fiber.Ctx) error {
	var url models.Url
	db := database.DBConn

	if err := c.BodyParser(&url); err != nil {
		return handlers.HandleError(c, http.StatusBadRequest, "Error parsing request body", err)
	}

	uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return handlers.HandleError(c, http.StatusInternalServerError, "Error uploading file", err)
	}

	// Insert video information into the database
	video := models.Video{
		Video_Url:  uploadUrl,
	}

	db.Create(&video)

	// Return the uploaded video information as JSON response
	return c.JSON(video)
}


func GetFiles(c *fiber.Ctx) error {
	db := database.DBConn
	var videos []models.Video
	db.Find(&videos)
	return c.JSON(videos)
}

func GetFile(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var video models.Video
	db.Find(&video, id)
	return c.JSON(video)
}

func DeleteFile(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var video models.Video
	db.First(&video, id)
	if video.Video_Url == "" {
		return handlers.HandleError(c, http.StatusBadRequest, "No Video found with given ID", fiber.ErrBadGateway)
	}
	db.Delete(&video)
	return handlers.SuccessResponse(c, http.StatusOK, &video)
}