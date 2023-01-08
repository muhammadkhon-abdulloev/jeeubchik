package account

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}
}
