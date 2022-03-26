package adapters

import (
	"github.com/fuato1/shorturl/internal/adapters/storage/redis"
	"github.com/fuato1/shorturl/internal/domain"
)

type Services struct {
	UrlSqlRepo   domain.UrlSqlRepository
	UrlCacheRepo domain.UrlCacheRepository
	QrSqlRepo    domain.QrRepository
}

func NewServices() *Services {
	return &Services{
		UrlSqlRepo:   nil,
		UrlCacheRepo: redis.NewRepo(),
		QrSqlRepo:    nil,
	}
}
