package queries

import (
	"log"

	"github.com/fuato1/shorturl/url-service/ports"
)

type GetAllUrlsResult struct {
	Urls map[string]string
}

type GetAllUrlsRequestHandler interface {
	Handle() (GetAllUrlsResult, error)
}

type getAllUrlsRequestHandler struct {
	repo ports.UrlRepository
}

func NewGetAllUrlsRequestHandler(repo ports.UrlRepository) *getAllUrlsRequestHandler {
	return &getAllUrlsRequestHandler{repo: repo}
}

func (rh getAllUrlsRequestHandler) Handle() (GetAllUrlsResult, error) {
	urls, err := rh.repo.GetAll()
	if err != nil {
		log.Fatalf("unable to GetAll Urls: %v", err)
		return GetAllUrlsResult{}, err
	}

	results := GetAllUrlsResult{
		Urls: urls,
	}

	return results, nil
}
