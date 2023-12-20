package handlers

import (
	"uploadService/mod/dtos"
	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, statusCode int, message string, err error) error {
	return c.Status(statusCode).JSON(dtos.MediaDto{
		StatusCode: statusCode,
		Message:    "error",
		Data:       &fiber.Map{"data": err.Error()},
	})
}

func SuccessResponse(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(dtos.MediaDto{
		StatusCode: statusCode,
		Message:    "success",
		Data:       &fiber.Map{"data": data},
	})
}