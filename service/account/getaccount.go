package account

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *Service) GetAccountByID(id uuid.UUID) (*Account, error) {
	account, err := s.accountRepo.GetAccountByID(id)
	if err != nil {
		return nil, fmt.Errorf("s.accountRepo.GetAccountByID: %w", err)
	}

	if account == nil {
		return nil, nil
	}

	account.UtilizePassword()

	return &Account{
		ID:        account.ID,
		Login:     account.Login,
		CreatedAt: account.CreatedAt,
	}, nil
}

type Account struct {
	ID        uuid.UUID
	Login     string
	CreatedAt time.Time
}
