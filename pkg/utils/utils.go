package utils

import (
	"contactsList/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

func CreateSessionCookie(c echo.Context, cfg *config.Cookie, session string) {
	c.SetCookie(&http.Cookie{
		Name:       cfg.Name,
		Value:      session,
		Path:       "/",
		RawExpires: "",
		MaxAge:     int(cfg.Expire),
		Secure:     cfg.Secure,
		HttpOnly:   cfg.HTTPOnly,
		SameSite:   0,
	})
}

func DeleteSessionCookie(c echo.Context, cookieName string) {
	c.SetCookie(&http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
