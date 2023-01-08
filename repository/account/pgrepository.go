package account

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

func (p *PGRepository) CreateAccount(account *Account) error {
	if account == nil {
		return errNilData
	}
	const query = `INSERT INTO account.account (id, login, password_hash) VALUES ($1, $2, $3);`
	_, err := p.conn.Exec(query, account.ID, account.Login, account.PasswordHash)
	if err != nil {
		return fmt.Errorf("p.conn.Exec: %w", err)
	}
	return nil
}

func (p *PGRepository) GetAccountByLogin(login string) (*Account, error) {
	var account Account
	const query = `SELECT id, login, password_hash, created_at, updated_at FROM account.account WHERE login = $1;`
	err := p.conn.Get(&account, query, login)
	if err != nil {
		return nil, fmt.Errorf("p.conn.Get: %w", err)
	}
	return &account, nil
}

func (p *PGRepository) GetAccountByID(id uuid.UUID) (*Account, error) {
	var account Account
	const query = `SELECT id, login, password_hash, created_at, updated_at FROM account.account WHERE id = $1;`
	err := p.conn.Get(&account, query, id)
	if err != nil {
		return nil, fmt.Errorf("p.conn.Get: %w", err)
	}
	return &account, nil
}

func (p *PGRepository) DeleteAccountByID(id uuid.UUID) error {
	const query = `DELETE FROM account.account WHERE id = $1;`
	_, err := p.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("p.conn.Exec: %w", err)
	}
	return nil
}
