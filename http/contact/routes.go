package contact

import (
	"contactsList/http/middleware"
	"github.com/labstack/echo/v4"
)

func MapContactRoutes(accountGroup *echo.Group, h *Handler, mw *middleware.Middleware) {
	accountGroup.GET("", h.GetAll(), mw.CookieMiddleware())
	accountGroup.POST("", h.CreateContact(), mw.CookieMiddleware())
	accountGroup.GET("/:id", h.GetContactByID(), mw.CookieMiddleware())
	accountGroup.PUT("/:id", h.UpdateContactByID(), mw.CookieMiddleware())
	accountGroup.DELETE("/:id", h.DeleteContactByID(), mw.CookieMiddleware())

}
