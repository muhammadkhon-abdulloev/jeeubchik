package account

import "github.com/google/uuid"

type Repository interface {
	create
	read
	update
	del
}

type create interface {
	CreateAccount(account *Account) error
}

type read interface {
	GetAccountByLogin(login string) (*Account, error)
	GetAccountByID(id uuid.UUID) (*Account, error)
}

type update interface {
}

type del interface {
	DeleteAccountByID(id uuid.UUID) error
}
