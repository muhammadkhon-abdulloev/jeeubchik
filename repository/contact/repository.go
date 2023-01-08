package contact

import "github.com/google/uuid"

type Repository interface {
	create
	read
	update
	del
}

type create interface {
	CreateContact(contact *Contact) error
}

type read interface {
	GetContactsByAccountID(id uuid.UUID) ([]Contact, error)
	GetContactByID(id, accountID uuid.UUID) (*Contact, error)
}

type update interface {
	UpdateContactByID(newData Contact, id uuid.UUID) error
}

type del interface {
	DeleteContactByID(id, contactID uuid.UUID) error
}
