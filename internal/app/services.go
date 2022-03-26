package app

import (
	"github.com/fuato1/shorturl/internal/app/qr"
	"github.com/fuato1/shorturl/internal/app/url"
)

// main service
type Services struct {
	UrlServices url.UrlServices
	QrServices  qr.QrServices
}

func NewServices(UrlServices url.UrlServices, QrServices qr.QrServices) Services {
	return Services{
		UrlServices: UrlServices,
		QrServices:  QrServices,
	}
}
