package http

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"` 
	Data    interface{} `json:"data,omitempty"`  
}

func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, errors interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		Status:  false,
		Message: message,
		Errors:  errors,
	})
}