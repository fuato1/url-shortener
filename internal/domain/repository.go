package domain

import (
	"github.com/fuato1/shorturl/internal/domain/qr"
	"github.com/fuato1/shorturl/internal/domain/url"
)

type QRRepository interface {
	GetAll() ([]qr.QR, error)
	Add(qr.QR) error
}

type URLRepository interface {
	GetAll() ([]url.ShortUrl, error)
	Add(url.ShortUrl) error
}
