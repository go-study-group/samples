package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/aaron/kv/pkg/database"
	"github.com/labstack/echo"
)

// Put is the handler for PUT /:key
func Put(c echo.Context) error {
	key := c.Param("key")
	req := c.Request()
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return c.String(400, "bad request")
	}
	val := string(b)
	database.Put(key, val)

	return c.String(http.StatusOK, "")
}
