package account

import (
	"contactsList/config"
	"contactsList/pkg/logger"
	"contactsList/pkg/utils"
	"contactsList/service/account"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *Handler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var params LoginParams
		if err := utils.ReadRequest(c, &params); err != nil {
			logger.GetLogger().Error(fmt.Errorf("utils.ReadRequest: %w", err))

			return c.NoContent(http.StatusBadRequest)
		}

		accountID, err := h.accountService.Login(account.LoginParams{Login: params.Login, Password: params.Password})
		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.accountService.Login: %w", err))

			return err
		}

		exp := config.GetConfig().Cookie.Expire * time.Hour
		sessionKey, err := h.sessionService.CreateSession(*accountID, exp)
		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.sessionService.CreateSession: %w", err))

			return c.NoContent(http.StatusInternalServerError)
		}

		utils.CreateSessionCookie(c, config.GetConfig().Cookie, sessionKey)

		return c.NoContent(http.StatusOK)
	}
}

type LoginParams struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
