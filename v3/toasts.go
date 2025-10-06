package htmx

import (
	"encoding/json"
	"errors"

	"github.com/katallaxie/htmx"

	"github.com/gofiber/fiber/v3"
)

// ErrToastsUnexpectedError is returned when there is an unexpected error.
var ErrToastsUnexpectedError = Error("there has been an unexpected error")

// ToastsErrorHandler is the default error handler for toasts.
var ToastsErrorHandler = func(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var toast Toast
	if !errors.As(err, &toast) {
		toast = ErrToastsUnexpectedError
	}

	var e *fiber.Error // if this is not a toast then use the error message
	if errors.As(err, &e) {
		code = e.Code
		toast = Error(e.Message)
	}

	if toast.Level != SUCCESS {
		ReSwap(c, "none")
	}

	err = toast.SetHXTriggerHeader(c)
	if err != nil {
		return err
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return c.Status(code).SendString(err.Error())
}

const (
	// INFO is the info level.
	INFO = "info"
	// SUCCESS is the success level.
	SUCCESS = "success"
	// ERROR is the error level.
	ERROR = "error"
)

// Toast is a message to display to the user.
type Toast struct {
	// Level is the level of the toast.
	Level string `json:"level"`
	// Message is the message of the toast.
	Message string `json:"message"`
	// Code is the http status code of the toast.
	Code int `json:"code"`
}

// New returns a new Toast.
func New(level string, message string, code ...int) Toast {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	return Toast{level, message, statusCode}
}

// Error returns the error message.
func (t Toast) Error() string {
	return t.Message
}

// Info returns an info message.
func Info(message string, code ...int) Toast {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	return New(INFO, message, statusCode)
}

// Success returns a success message.
func Success(c fiber.Ctx, message string, code ...int) {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	//nolint:errcheck
	New(SUCCESS, message, statusCode).SetHXTriggerHeader(c)
}

// Error returns an error message.
func Error(message string, code ...int) Toast {
	statusCode := 500
	if len(code) > 0 {
		statusCode = code[0]
	}

	return New(ERROR, message, statusCode)
}

// ToJSON returns the JSON representation of the toast.
func (t Toast) ToJSON() (string, error) {
	t.Message = t.Error()

	eventMap := map[string]Toast{}
	eventMap["htmx-toasts:notify"] = t

	jsonData, err := json.Marshal(eventMap)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// SetHXTriggerHeader sets the HTMX trigger header.
func (t Toast) SetHXTriggerHeader(c fiber.Ctx) error {
	jsonData, err := t.ToJSON()
	if err != nil {
		return err
	}

	Trigger(c, jsonData)

	return nil
}

// Toast is the toast component.
func Toasts() htmx.Node {
	return htmx.CustomElement("htmx-toasts")
}
