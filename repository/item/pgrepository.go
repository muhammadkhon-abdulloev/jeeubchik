package item

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PGRepository struct {
	conn *sqlx.DB
}

var _ Repository = (*PGRepository)(nil)

func NewPGRepository(conn *sqlx.DB) *PGRepository {
	return &PGRepository{
		conn: conn,
	}
}

func (p *PGRepository) CreateItem(account *Item) error {
	//TODO implement me
	panic("implement me")
}

func (p *PGRepository) GetItemByContactID(id uuid.UUID) (*Item, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PGRepository) GetItemByID(id uuid.UUID) (*Item, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PGRepository) DeleteItemByID(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
