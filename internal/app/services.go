package app

import (
	"github.com/fuato1/shorturl/internal/app/qr"
	"github.com/fuato1/shorturl/internal/app/url"
)

// main service
type Services struct {
	URLServices url.URLServices
	QRServices  qr.QRServices
}

func NewServices(URLServices url.URLServices, QRServices qr.QRServices) Services {
	return Services{
		URLServices: URLServices,
		QRServices:  QRServices,
	}
}
