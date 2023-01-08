package account

import (
	"contactsList/config"
	"contactsList/pkg/logger"
	"contactsList/pkg/utils"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(config.GetConfig().Cookie.Name)
		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("c.Cookie: %w", err))

			return c.NoContent(http.StatusUnauthorized)
		}

		if cookie.Value == "" {
			return c.NoContent(http.StatusUnauthorized)
		}
		err = h.sessionService.DeleteSessionByKey(context.Background(), cookie.Value)
		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.sessUC.DeleteSessionByKey: %w", err))

			return c.NoContent(http.StatusUnauthorized)
		}
		utils.DeleteSessionCookie(c, config.GetConfig().Cookie.Name)

		return c.NoContent(http.StatusOK)
	}
}
