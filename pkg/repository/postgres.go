package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	linksTable = "links"
)

type Config struct {
	URI string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.URI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
