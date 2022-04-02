package ports

import (
	"github.com/fuato1/shorturl/internal/app"
	"github.com/fuato1/shorturl/internal/ports/http/qr"
	"github.com/fuato1/shorturl/internal/ports/http/url"
)

type UrlServices struct {
	UrlServer *url.UrlServer
}

type QrServices struct {
	QrServer *qr.QrServer
}

func NewUrlServices(appServices app.UrlServices) UrlServices {
	return UrlServices{
		UrlServer: url.NewUrlServer(appServices.UrlServices),
	}
}

func NewQrServices(appServices app.QrServices) QrServices {
	return QrServices{
		QrServer: qr.NewQrServer(appServices.QrServices),
	}
}
