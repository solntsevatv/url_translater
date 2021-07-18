package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/solntsevatv/url_translater/internal/url_translater"
)

type Url interface {
	GetNextUrlId() (int, error)
	CreateShortURL(url url_translater.URL) (string, error)
	GetLongURL(short_url url_translater.ShortURL) (string, error)
}

type Repository struct {
	Url
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Url: NewUrlPostgres(db),
	}
}
