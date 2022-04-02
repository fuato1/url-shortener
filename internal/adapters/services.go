package adapters

import (
	"github.com/fuato1/shorturl/internal/adapters/storage/redis"
	"github.com/fuato1/shorturl/internal/domain"
)

type UrlServices struct {
	UrlSqlRepo   domain.UrlSqlRepository
	UrlCacheRepo domain.UrlCacheRepository
}

type QrServices struct {
	QrSqlRepo domain.QrRepository
}

func NewUrlServices() *UrlServices {
	return &UrlServices{
		UrlSqlRepo:   nil,
		UrlCacheRepo: redis.NewRepo(),
	}
}

func NewQrServices() *QrServices {
	return &QrServices{
		QrSqlRepo: nil,
	}
}
