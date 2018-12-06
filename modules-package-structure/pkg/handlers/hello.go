package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Hello is a simple handler for the '/' route
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
