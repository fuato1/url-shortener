package domain

import (
	"github.com/fuato1/shorturl/internal/domain/qr"
	"github.com/fuato1/shorturl/internal/domain/url"
)

// url sql repository
type UrlSqlRepository interface {
	GetAll() ([]url.ShortUrl, error)
	Add(url.ShortUrl) error
	Get(shortUrl string) (string, error)
}

// url cache repository
type UrlCacheRepository interface {
	GetAll() (map[string]string, error)
	Add(url url.ShortUrl) error
	Get(shortUrl string) (string, error)
}

// qr repository
type QrRepository interface {
	GetAll() ([]qr.QR, error)
	Add(qr.QR) error
}
