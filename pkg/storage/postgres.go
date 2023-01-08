package storage

import (
	"contactsList/config"
	"fmt"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
	"time"
)

func InitPsqlDB(c *config.Postgres) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode)

	database, err := sqlx.Connect(c.PgDriver, connectionUrl)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Connect: %w", err)
	}

	database.SetMaxOpenConns(c.Settings.MaxOpenConns)
	database.SetConnMaxLifetime(c.Settings.ConnMaxLifetime * time.Second)
	database.SetMaxIdleConns(c.Settings.MaxIdleConns)
	database.SetConnMaxIdleTime(c.Settings.ConnMaxIdleTime * time.Second)

	if err = database.Ping(); err != nil {
		return nil, fmt.Errorf("database.Ping: %w", err)
	}
	return database, nil
}
