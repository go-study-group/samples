package handlers

import (
	"fmt"
	"net/http"

	"github.com/aaron/kv/pkg/database"
	"github.com/labstack/echo"
)

// Get is the handler for GET /:key
func Get(c echo.Context) error {
	key := c.Param("key")
	val, ok := database.Get(key)
	if !ok {
		return c.String(404, "not found")
	}
	str := fmt.Sprintf("This is the Get! You passed in %s, got value %s", key, val)
	return c.String(http.StatusOK, str)
}
