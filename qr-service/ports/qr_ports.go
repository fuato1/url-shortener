package ports

import "github.com/fuato1/shorturl/qr-service/model"

// qr repository
type QrRepository interface {
	GetAll() ([]model.QR, error)
	Add(model.QR) error
}
