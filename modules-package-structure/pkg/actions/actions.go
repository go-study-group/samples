package actions

import (
	"github.com/aaron/kv/pkg/handlers"
	"github.com/labstack/echo"
)

// Apply applies routes for my server
func Apply(e *echo.Echo) {
	e.GET("/", handlers.Hello)
	e.GET("/:key", handlers.Get)
	e.PUT("/:key", handlers.Put)
}
