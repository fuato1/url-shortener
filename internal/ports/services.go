package ports

import (
	"github.com/fuato1/shorturl/internal/app"
	"github.com/fuato1/shorturl/internal/ports/http/qr"
	"github.com/fuato1/shorturl/internal/ports/http/url"
)

type Services struct {
	UrlServer *url.UrlServer
	QrServer  *qr.QrServer
}

func NewServices(appServices app.Services) Services {
	return Services{
		UrlServer: url.NewUrlServer(appServices.UrlServices),
		QrServer:  qr.NewQrServer(appServices.QrServices),
	}
}
