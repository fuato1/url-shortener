package queries

import (
	"log"

	"github.com/fuato1/shorturl/internal/domain"
	"github.com/fuato1/shorturl/internal/domain/url"
)

type GetAllURLsResult struct {
	URL url.URL
}

type GetAllURLsRequestHandler interface {
	Handle() ([]GetAllURLsResult, error)
}

type getAllURLsRequestHandler struct {
	repo domain.URLRepository
}

func NewGetAllURLsRequestHandler(repo domain.URLRepository) *getAllURLsRequestHandler {
	return &getAllURLsRequestHandler{repo: repo}
}

func (rh getAllURLsRequestHandler) Handle() ([]GetAllURLsResult, error) {
	urls, err := rh.repo.GetAll()
	if err != nil {
		log.Fatalf("unable to GetAll URLs: %v", err)
		return []GetAllURLsResult{}, err
	}

	var results []GetAllURLsResult
	for _, url := range urls {
		results = append(results, GetAllURLsResult{URL: url})
	}

	return results, nil
}
