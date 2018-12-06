package main

import (
	"github.com/aaron/kv/pkg/actions"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	actions.Apply(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
