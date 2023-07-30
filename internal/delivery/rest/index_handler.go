package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetHelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": map[string]interface{}{
			"code":    200,
			"message": "Hello World!",
		},
	})
}
