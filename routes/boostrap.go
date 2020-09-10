package routes

import (
	"github.com/labstack/echo/v4"
)

// Boostrap ...
func Boostrap(e *echo.Echo) {
	Transaction(e)
}
