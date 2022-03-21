package adapters

import (
	"github.com/fuato1/shorturl/internal/adapters/storage/redis"
	"github.com/fuato1/shorturl/internal/domain"
)

type Services struct {
	URLRepository domain.URLRepository
	QRRepository  domain.QRRepository
}

func NewServices() *Services {
	return &Services{
		URLRepository: redis.NewRepo(),
		// QrRepository: postgres.NewRepo(),
	}
}
