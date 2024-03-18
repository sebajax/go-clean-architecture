package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sebajax/go-clean-architecture/pkg/apperror"
)

// func to manage error response messages in middleware
func errorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

// this is a middleware that converts AppError to fiber.Error.
func ErrorHandler(c *fiber.Ctx) error {
	// try to execute the next middleware/handler
	err := c.Next()

	// check if there was an error
	if err != nil {
		// log the error, handle it, or send a custom response
		if e, ok := err.(*apperror.AppError); ok {
			log.Error(errorResponse(e))
			return c.Status(e.Code).JSON(errorResponse(e))
		}

		// internal server error ocurred trying to cast error to apperror
		log.Error(errorResponse(err))
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse(err))
	}

	// if no error, continue to execute
	return nil
}
