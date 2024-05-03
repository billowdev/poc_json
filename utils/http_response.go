package utils

import "github.com/gofiber/fiber/v2"

type SAPIResponse[T interface{}] struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Data          T      `json:"data"`
}

func NewResponse[T interface{}](c *fiber.Ctx, status_code string, message string, data T) error {
	response := SAPIResponse[T]{
		StatusCode:    status_code,
		StatusMessage: message,
		Data:          data,
	}
	return c.Status(200).JSON(response)
}
func NewSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	response := SAPIResponse[interface{}]{
		StatusCode:    "S-000",
		StatusMessage: message,
		Data:          data,
	}
	return c.Status(200).JSON(response)
}
