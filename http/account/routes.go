package account

import (
	"contactsList/http/middleware"
	"github.com/labstack/echo/v4"
)

func MapAccountRoutes(accountGroup *echo.Group, h *Handler, mw *middleware.Middleware) {
	accountGroup.POST("/login", h.Login())
	accountGroup.POST("/register", h.Register())
	accountGroup.POST("/logout", h.Logout(), mw.CookieMiddleware())
	accountGroup.GET("/account", h.GetAccount(), mw.CookieMiddleware())
}
