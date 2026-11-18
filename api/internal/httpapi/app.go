package httpapi

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

// bodyLimit sits above the 4 MB image rule (I2) so oversized uploads reach
// the handler and get our file_too_large error instead of Fiber's default 413.
const bodyLimit = 5 * 1024 * 1024

// NewApp wires the generated routes onto a Fiber app with the API's error format.
func NewApp(server StrictServerInterface) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "roadmap-api",
		BodyLimit:    bodyLimit,
		ErrorHandler: errorHandler,
	})
	RegisterHandlers(app, NewStrictHandler(server, nil))
	return app
}

// errorHandler renders every unhandled error as the contract's Error schema.
func errorHandler(c fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	var fe *fiber.Error
	if errors.As(err, &fe) {
		status = fe.Code
	}

	body := Error{Code: codeForStatus(status), Message: fiber.ErrInternalServerError.Message}
	if status < fiber.StatusInternalServerError || status == fiber.StatusNotImplemented {
		body.Message = err.Error()
	} else {
		slog.Error("request failed", "method", c.Method(), "path", c.Path(), "err", err)
	}
	return c.Status(status).JSON(body)
}

func codeForStatus(status int) ErrorCode {
	switch status {
	case fiber.StatusBadRequest:
		return ErrorCodeValidationFailed
	case fiber.StatusUnauthorized:
		return ErrorCodeUnauthorized
	case fiber.StatusForbidden:
		return ErrorCodeForbidden
	case fiber.StatusNotFound, fiber.StatusMethodNotAllowed:
		return ErrorCodeNotFound
	case fiber.StatusRequestEntityTooLarge:
		return ErrorCodeFileTooLarge
	case fiber.StatusUnsupportedMediaType:
		return ErrorCodeUnsupportedMediaType
	default:
		return ErrorCodeInternalError
	}
}
