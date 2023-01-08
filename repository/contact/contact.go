package contact

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	errNilData = errors.New("nil parameter got")
)

type Contact struct {
	ID        uuid.UUID `db:"id" json:"ID"`
	AccountID uuid.UUID `db:"account_id" json:"accountID"`
	FullName  string    `db:"full_name" json:"fullName"`
	Email     string    `db:"email" json:"email"`
	Phone     string    `db:"phone" json:"phone"`
	Address   string    `db:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
