package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"uploadService/mod/controllers"
	"uploadService/mod/database"
	"uploadService/mod/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFileUpload(t *testing.T) {
	// Initialize Fiber app and set up database connection
	app := fiber.New()
	database.InitDatabase()

	// Create a test request with a sample file
	request := httptest.NewRequest(http.MethodPost, "/file-upload", nil)
	request.Header.Set("Content-Type", "multipart/form-data")
	request.Header.Set("Content-Disposition", "form-data; name=\"file\"; filename=\"testfile.txt\"")

	// Create a test context with the request
	ctx := app.AcquireCtx(request)
	defer app.ReleaseCtx(ctx)

	// Call the FileUpload function
	err := controllers.FileUpload(ctx)
	assert.Nil(t, err)

	// Check the response status code
	assert.Equal(t, http.StatusOK, ctx.Response().StatusCode())
}

func TestRemoteUpload(t *testing.T) {
	// Initialize Fiber app and set up database connection
	app := fiber.New()
	database.InitDatabase()

	// Create a test request with a JSON body
	url := models.Url{Url: "https://example.com/video.mp4"}
	body, _ := json.Marshal(url)
	request := httptest.NewRequest(http.MethodPost, "/remote-upload", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	// Create a test context with the request
	ctx := app.AcquireCtx(request)
	defer app.ReleaseCtx(ctx)

	// Call the RemoteUpload function
	err := controllers.RemoteUpload(ctx)
	assert.Nil(t, err)

	// Check the response status code
	assert.Equal(t, http.StatusOK, ctx.Response().StatusCode())
}

func TestGetFiles(t *testing.T) {
	// Initialize Fiber app and set up database connection
	app := fiber.New()
	database.InitDatabase()

	// Create a test request for GetFiles
	request := httptest.NewRequest(http.MethodGet, "/get-files", nil)

	// Create a test context with the request
	ctx := app.AcquireCtx(request, nil)
	defer app.ReleaseCtx(ctx)

	// Call the GetFiles function
	err := controllers.GetFiles(ctx)
	assert.Nil(t, err)

	// Check the response status code
	assert.Equal(t, http.StatusOK, ctx.Response().StatusCode())
}

func TestGetFile(t *testing.T) {
	// Initialize Fiber app and set up database connection
	app := fiber.New()
	database.InitDatabase()

	// Create a test video entry in the database
	db := database.DBConn
	video := models.Video{Video_Url: "https://example.com/video.mp4"}
	db.Create(&video)

	// Create a test request for GetFile
	request := httptest.NewRequest(http.MethodGet, "/get-file/"+video.ID, nil)

	// Create a test context with the request
	ctx := app.AcquireCtx(request, nil)
	defer app.ReleaseCtx(ctx)

	// Call the GetFile function
	err := controllers.GetFile(ctx)
	assert.Nil(t, err)

	// Check the response status code
	assert.Equal(t, http.StatusOK, ctx.Response().StatusCode())
}

func TestDeleteFile(t *testing.T) {
	// Initialize Fiber app and set up database connection
	app := fiber.New()
	database.InitDatabase()

	// Create a test video entry in the database
	db := database.DBConn
	video := models.Video{Video_Url: "https://example.com/video.mp4"}
	db.Create(&video)

	// Create a test request for DeleteFile
	request := httptest.NewRequest(http.MethodDelete, "/delete-file/"+video.ID, nil)

	// Create a test context with the request
	ctx := app.AcquireCtx(request, nil)
	defer app.ReleaseCtx(ctx)

	// Call the DeleteFile function
	err := controllers.DeleteFile(ctx)
	assert.Nil(t, err)

	// Check the response status code
	assert.Equal(t, http.StatusOK, ctx.Response().StatusCode())
}
