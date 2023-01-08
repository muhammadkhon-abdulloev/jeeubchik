package item

import "github.com/google/uuid"

type Repository interface {
	create
	read
	update
	del
}

type create interface {
	CreateItem(account *Item) error
}

type read interface {
	GetItemByContactID(id uuid.UUID) (*Item, error)
	GetItemByID(id uuid.UUID) (*Item, error)
}

type update interface {
}

type del interface {
	DeleteItemByID(id uuid.UUID) error
}
