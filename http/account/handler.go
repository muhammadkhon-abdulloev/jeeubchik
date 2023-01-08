package account

import (
	"contactsList/service/account"
	"contactsList/service/session"
)

type Handler struct {
	accountService *account.Service
	sessionService *session.Service
}

func NewHandler(accountService *account.Service, sessionService *session.Service) *Handler {
	return &Handler{
		accountService: accountService,
		sessionService: sessionService,
	}
}
