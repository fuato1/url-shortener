package ports

import "github.com/fuato1/shorturl/url-service/model"

type UrlRepository interface {
	GetAll() (map[string]string, error)
	Add(url model.ShortUrl) error
	Get(shortUrl string) (string, error)
}
