package account

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	errNilData = errors.New("nil parameter got")
)

type Account struct {
	ID           uuid.UUID `db:"id"`
	Login        string    `db:"login"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (a *Account) UtilizePassword() {
	a.PasswordHash = ""
}
