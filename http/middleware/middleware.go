package middleware

import (
	"contactsList/service/account"
	"contactsList/service/session"
)

type Middleware struct {
	accountService *account.Service
	sessionService *session.Service
}

func NewMiddleware(
	accountService *account.Service,
	sessionService *session.Service,
) *Middleware {
	return &Middleware{
		accountService: accountService,
		sessionService: sessionService,
	}
}
