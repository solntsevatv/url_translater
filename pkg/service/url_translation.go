package service

import (
	"github.com/solntsevatv/url_translater/internal/url_translater"
	"github.com/solntsevatv/url_translater/pkg/repository"
)

const digits = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

type UrlTranslationService struct {
	repo repository.Url
}

func NewUrlTranslationService(repo repository.Url) *UrlTranslationService {
	return &UrlTranslationService{repo: repo}
}

func (s *UrlTranslationService) GetNextUrlId() (int, error) {
	return s.repo.GetNextUrlId()
}

func (s *UrlTranslationService) CreateShortURL(long_url url_translater.LongURL) (string, error) {
	url_id, err := s.GetNextUrlId()
	if err != nil {
		return "", err
	}
	long_url.Id = url_id

	short_url := TranslateLongToShort(long_url)

	return s.repo.CreateShortURL(url_translater.URL{
		Id:       long_url.Id,
		LongUrl:  long_url.LinkUrl,
		ShortURL: short_url,
	})
}

func TranslateLongToShort(long_url url_translater.LongURL) string {
	short_url := ""
	bytes := []byte(digits)
	for id := long_url.Id - 1; id >= 0; id = int(float64(id)/62) - 1 {
		ind := id % 62
		short_url = string(bytes[ind]) + short_url
	}
	return short_url
}

func (s *UrlTranslationService) GetLongURL(short_url url_translater.ShortURL) (string, error) {
	return s.repo.GetLongURL(short_url)
}
