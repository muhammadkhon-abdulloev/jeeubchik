package account

import (
	"contactsList/repository/account"
)

type Service struct {
	accountRepo account.Repository
}

func NewService(
	accountRepo account.Repository,
) *Service {
	return &Service{
		accountRepo: accountRepo,
	}
}
