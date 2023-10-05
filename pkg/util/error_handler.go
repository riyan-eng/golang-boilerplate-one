package util

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var e *fiber.Error
	code := fiber.StatusInternalServerError
	if errors.As(err, &e) {
		code = e.Code
		c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		return c.Status(code).SendString(err.Error())
	}
	// logger
	log.Println(err)

	// validation
	_, ok := err.(BadRequest)
	if ok {
		return NewResponse(c).Error(err.Error(), MESSAGE_BAD_REQUEST, fiber.StatusBadRequest)
	}

	_, ok = err.(ValidationError)
	if ok {

		return NewResponse(c).Error(err.Error(), MESSAGE_FAILED_VALIDATION, fiber.StatusBadRequest)
	}

	// no data
	if err == sql.ErrNoRows {
		return NewResponse(c).Error(err.Error(), MESSAGE_NOT_FOUND, fiber.StatusBadRequest)
	}

	return NewResponse(c).Error(err.Error(), MESSAGE_BAD_SYSTEM, fiber.StatusBadGateway)
}
