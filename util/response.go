package util

import (
	"github.com/labstack/echo/v4"

	"net/http"
)

// Response ...
type Response map[string]interface{}

func generateResponse(data interface{}, message string) Response {
	return Response{
		"data":    data,
		"message": message,
	}
}

// Response200 success.....
func Response200(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Thanh Cong!"
	}
	return c.JSON(http.StatusOK, generateResponse(data, message))
}

// Response400 badrequest ...
func Response400(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Du lieu khong hop le"
	}
	return c.JSON(http.StatusBadRequest, generateResponse(data, message))
}

// Response404 not found
func Response404(c echo.Context, data interface{}, message string) error {
	if message == "" {
		message = "Du lieu khong tim thay"
	}
	return c.JSON(http.StatusNotFound, generateResponse(data, message))
}

// Response 
