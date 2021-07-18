package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/solntsevatv/url_translater/internal/url_translater"
)

type UrlPostgres struct {
	db *sqlx.DB
}

func NewUrlPostgres(db *sqlx.DB) *UrlPostgres {
	return &UrlPostgres{db: db}
}

func (r *UrlPostgres) IsDBEmpty() (bool, error) {
	var count int
	query := fmt.Sprintf("SELECT count(*) FROM %s;", linksTable)
	err := r.db.Get(&count, query)
	if err != nil {
		return true, err
	}

	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *UrlPostgres) GetNextUrlId() (int, error) {
	var id int

	is_empty, err := r.IsDBEmpty()
	if err != nil {
		return 0, err
	}
	if is_empty {
		id = 1
		return id, nil
	}

	query := fmt.Sprintf("SELECT last_value FROM %s_id_seq;", linksTable)
	err = r.db.Get(&id, query)
	if err != nil {
		return 0, err
	}
	return id + 1, nil // прибавляем 1, так как возвращаем СЛЕДУЮЩИЙ id
}

func (r *UrlPostgres) CreateShortURL(url url_translater.URL) (string, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (long_url, short_url) values ($1, $2) RETURNING id", linksTable)
	row := r.db.QueryRow(query, url.LongUrl, url.ShortURL)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return url.ShortURL, nil
}

func (r *UrlPostgres) GetLongURL(short_url url_translater.ShortURL) (string, error) {
	var long_url url_translater.LongURL
	query := fmt.Sprintf("SELECT id, long_url FROM %s WHERE short_url=$1;", linksTable)

	err := r.db.Get(&long_url, query, short_url.LinkUrl)
	if err != nil {
		return "", err
	}
	return long_url.LinkUrl, nil
}
