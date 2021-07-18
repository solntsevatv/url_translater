package service

import (
	"github.com/solntsevatv/url_translater/internal/url_translater"
	"github.com/solntsevatv/url_translater/pkg/repository"
)

//go:generate mockgen -destination=pkg/service/mocks_pkg/service.go  github.com/solntsevatv/url_translater/pkg/service UrlTranslation

type UrlTranslation interface {
	GetNextUrlId() (int, error)
	CreateShortURL(long_url url_translater.LongURL) (string, error)
	GetLongURL(short_url url_translater.ShortURL) (string, error)
}

type Service struct {
	UrlTranslation
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UrlTranslation: NewUrlTranslationService(repos.Url),
	}
}
