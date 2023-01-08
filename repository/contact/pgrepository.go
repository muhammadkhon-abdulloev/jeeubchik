package contact

import (
	"fmt"
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

func (p *PGRepository) CreateContact(contact *Contact) error {
	if contact == nil {
		return errNilData
	}
	const query = `INSERT INTO account.contact (id, account_id, full_name, email, phone, address) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := p.conn.Exec(query, contact.ID, contact.AccountID, contact.FullName, contact.Email, contact.Phone, contact.Address)
	if err != nil {
		return fmt.Errorf("p.conn.Exec: %w", err)
	}
	return nil
}

func (p *PGRepository) GetContactsByAccountID(id uuid.UUID) ([]Contact, error) {
	var accounts = make([]Contact, 0)
	const query = `SELECT id, account_id, full_name, email, phone, address, created_at, updated_at FROM account.contact WHERE account_id = $1;`
	err := p.conn.Select(&accounts, query, id)
	if err != nil {
		return nil, fmt.Errorf("p.conn.Select: %w", err)
	}
	return accounts, nil
}

func (p *PGRepository) GetContactByID(id, accountID uuid.UUID) (*Contact, error) {
	var account Contact
	const query = `SELECT id, account_id, full_name, email, phone, address, created_at, updated_at FROM account.contact WHERE id = $1 and account_id = $2;`
	err := p.conn.Get(&account, query, id, accountID)
	if err != nil {
		return nil, fmt.Errorf("p.conn.Get: %w", err)
	}
	return &account, nil
}

func (p *PGRepository) UpdateContactByID(newData Contact, id uuid.UUID) error {
	const query = `UPDATE account.contact set full_name = $2, email = $3, phone = $4, address = $5 WHERE id = $1;`
	_, err := p.conn.Exec(query, id, newData.FullName, newData.Email, newData.Phone, newData.Address)
	if err != nil {
		return fmt.Errorf("p.conn.Exec: %w", err)
	}
	return nil
}

func (p *PGRepository) DeleteContactByID(id, contactID uuid.UUID) error {
	const query = `DELETE FROM account.contact WHERE id = $1 and account_id = $2;`
	_, err := p.conn.Exec(query, id, contactID)
	if err != nil {
		return fmt.Errorf("p.conn.Exec: %w", err)
	}
	return nil
}
