package app

import (
	"github.com/fuato1/shorturl/internal/app/qr"
	"github.com/fuato1/shorturl/internal/app/url"
)

type UrlServices struct {
	UrlServices url.UrlServices
}

type QrServices struct {
	QrServices qr.QrServices
}

func NewUrlServices(urlServices url.UrlServices) UrlServices {
	return UrlServices{
		UrlServices: urlServices,
	}
}

func NewQrServices(qrServices qr.QrServices) QrServices {
	return QrServices{
		QrServices: qrServices,
	}
}
