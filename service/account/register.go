package account

import (
	"contactsList/repository/account"
	"fmt"
	"github.com/google/uuid"
)

func (s *Service) Register(register AccountRegister) (uuid.UUID, error) {
	hashedPass, err := HashPassword(register.Password)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("HashPassword: %w", err)
	}

	data := &account.Account{
		ID:           uuid.New(),
		Login:        register.Login,
		PasswordHash: hashedPass,
	}
	err = s.accountRepo.CreateAccount(data)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("s.accountRepo.CreateAccount: %w", err)

	}
	return data.ID, nil
}

type AccountRegister struct {
	Login    string
	Password string
}
