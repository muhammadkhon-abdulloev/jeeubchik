package account

import (
	"contactsList/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func (h *Handler) GetAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID, ok := c.Get("accountID").(uuid.UUID)
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}

		account, err := h.accountService.GetAccountByID(accountID)
		if err != nil {
			logger.GetLogger().Error("h.accountService.GetAccountByID: %w", err)

			return err
		}
		return c.JSON(http.StatusOK, Account{
			ID:        account.ID,
			Login:     account.Login,
			CreatedAt: account.CreatedAt,
		})
	}
}

type Account struct {
	ID        uuid.UUID `json:"id"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"createdAt"`
}
