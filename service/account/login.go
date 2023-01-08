package account

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func (s *Service) Login(params LoginParams) (*uuid.UUID, error) {
	foundAccount, err := s.accountRepo.GetAccountByLogin(params.Login)
	if err != nil {
		return nil, fmt.Errorf("s.accountRepo.GetAccountByLogin: %w", err)
	}

	if foundAccount == nil {
		return nil, errors.New("invalid data")
	}

	if err = ComparePasswords(foundAccount.PasswordHash, params.Password); err != nil {
		return nil, errors.New("invalid data")
	}

	return &foundAccount.ID, nil
}

type LoginParams struct {
	Login    string
	Password string
}
