package contact

import (
	"contactsList/pkg/logger"
	"contactsList/pkg/utils"
	"contactsList/repository/contact"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	contactRepo contact.Repository
}

func NewHandler(contactService contact.Repository) *Handler {
	return &Handler{
		contactRepo: contactService,
	}
}

func (h *Handler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID, ok := c.Get("accountID").(uuid.UUID)
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}

		accounts, err := h.contactRepo.GetContactsByAccountID(accountID)
		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.contactRepo.GetContactsByAccountID: %w", err))
			if errors.Is(err, sql.ErrNoRows) {
				return c.NoContent(http.StatusNoContent)
			}
			return err
		}

		return c.JSON(http.StatusOK, accounts)
	}
}

func (h *Handler) CreateContact() echo.HandlerFunc {
	type Contact struct {
		FullName string `json:"FullName" validate:"required"`
		Email    string `json:"Email" validate:"required"`
		Phone    string `json:"Phone" validate:"required"`
		Address  string `json:"Address" validate:"required"`
	}
	return func(c echo.Context) error {
		accountID, ok := c.Get("accountID").(uuid.UUID)
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}

		var params Contact
		if err := utils.ReadRequest(c, &params); err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.contactRepo.GetContactsByAccountID: %w", err))
			return c.NoContent(http.StatusBadRequest)
		}

		err := h.contactRepo.CreateContact(&contact.Contact{
			ID:        uuid.New(),
			AccountID: accountID,
			FullName:  params.FullName,
			Email:     params.Email,
			Phone:     params.Phone,
			Address:   params.Address,
		})

		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.contactRepo.CreateContact: %w", err))
			if errors.Is(err, sql.ErrNoRows) {
				return c.NoContent(http.StatusNoContent)
			}
			return err
		}

		return c.NoContent(http.StatusOK)

	}
}

func (h *Handler) GetContactByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID, ok := c.Get("accountID").(uuid.UUID)
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}
		id := c.Param("id")

		contactID, err := uuid.Parse(id)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		foundContact, err := h.contactRepo.GetContactByID(contactID, accountID)
		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.contactRepo.CreateContact: %w", err))
			if errors.Is(err, sql.ErrNoRows) {
				return c.NoContent(http.StatusNoContent)
			}
			return err
		}

		return c.JSON(http.StatusOK, foundContact)
	}
}

func (h *Handler) UpdateContactByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusUnauthorized)
	}
}

func (h *Handler) DeleteContactByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountID, ok := c.Get("accountID").(uuid.UUID)
		if !ok {
			return c.NoContent(http.StatusUnauthorized)
		}
		id := c.Param("id")

		contactID, err := uuid.Parse(id)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err = h.contactRepo.DeleteContactByID(contactID, accountID)
		if err != nil {
			logger.GetLogger().Error(fmt.Errorf("h.contactRepo.DeleteContactByID: %w", err))
			if errors.Is(err, sql.ErrNoRows) {
				return c.NoContent(http.StatusBadRequest)
			}
			return err
		}

		return c.NoContent(http.StatusOK)

	}
}
