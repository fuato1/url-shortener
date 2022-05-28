package commands

import (
	"log"

	"github.com/fuato1/shorturl/pkg/shortener"
	"github.com/fuato1/shorturl/pkg/time"
	"github.com/fuato1/shorturl/pkg/uuid"
	"github.com/fuato1/shorturl/url-service/model"
	"github.com/fuato1/shorturl/url-service/ports"
)

type AddUrlRequest struct {
	Source string
	UserId string
}

type AddUrlRequestHandler interface {
	Handle(req AddUrlRequest) error
}

type addUrlRequestHandler struct {
	repo       ports.UrlRepository
	uuidP      uuid.Provider
	timeP      time.Provider
	shortenerP shortener.Provider
}

func NewAddUrlRequestHandler(repo ports.UrlRepository, uuidP uuid.Provider, timeP time.Provider, shortenerP shortener.Provider) *addUrlRequestHandler {
	return &addUrlRequestHandler{
		repo:       repo,
		uuidP:      uuidP,
		timeP:      timeP,
		shortenerP: shortenerP,
	}
}

func (rh addUrlRequestHandler) Handle(req AddUrlRequest) error {
	url := model.ShortUrl{
		Id:        rh.uuidP.NewUUID(),
		Source:    req.Source,
		URL:       rh.shortenerP.Shorten(req.Source, req.UserId),
		CreatedAt: rh.timeP.Now(),
	}

	err := rh.repo.Add(url)
	if err != nil {
		log.Fatalf("unable to add shortened url to repository: %v", err)
		return err
	}

	return nil
}
