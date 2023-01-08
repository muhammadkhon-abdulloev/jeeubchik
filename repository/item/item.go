package item

import (
	"errors"
	"github.com/google/uuid"
)

var (
	errNilData = errors.New("nil parameter got")
)

type Item struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       float64   `db:"price"`
	ContactID   uuid.UUID `db:"contact_id"`
}
