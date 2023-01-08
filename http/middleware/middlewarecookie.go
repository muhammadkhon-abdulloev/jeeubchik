package middleware

import (
	"contactsList/config"
	"contactsList/pkg/logger"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (m *Middleware) CookieMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(config.GetConfig().Cookie.Name)
			if err != nil {
				logger.GetLogger().Error(fmt.Errorf("c.Cookie: %w", err))
				return c.NoContent(http.StatusUnauthorized)
			}
			if cookie.Value == "" {
				return c.NoContent(http.StatusUnauthorized)
			}
			accountID, err := m.sessionService.GetSession(c.Request().Context(), cookie.Value)
			if err != nil && !errors.Is(err, redis.Nil) {
				logger.GetLogger().Error(fmt.Errorf("m.sessionService.GetSession: %w", err))
				return c.NoContent(http.StatusInternalServerError)
			}

			if accountID == nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			account, err := m.accountService.GetAccountByID(*accountID)
			if err != nil {
				logger.GetLogger().Error(fmt.Errorf("m.accountService.GetAccountByID: %w", err))
				return c.NoContent(http.StatusUnauthorized)
			}
			if account == nil {
				return c.NoContent(http.StatusUnauthorized)
			}

			c.Set("accountID", account.ID)
			return next(c)
		}
	}
}
