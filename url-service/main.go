package main

import (
	"github.com/fuato1/shorturl/pkg/base58"
	"github.com/fuato1/shorturl/pkg/sha256"
	"github.com/fuato1/shorturl/pkg/shortener"
	"github.com/fuato1/shorturl/pkg/time"
	"github.com/fuato1/shorturl/pkg/uuid"
	"github.com/fuato1/shorturl/url-service/handlers"
	"github.com/fuato1/shorturl/url-service/repository"
	"github.com/fuato1/shorturl/url-service/services"
)

func main() {
	// creating adapter services
	repo := repository.NewUrlRepository()

	// creating providers
	uuidP := uuid.NewUUIDProvider()
	timeP := time.NewTimeProvider()
	shortenerP := shortener.NewShortenerProvider(
		sha256.NewSHA256Provider(),
		base58.NewBase58Provider(),
	)

	// injecting dependencies
	urlServices := services.NewUrlServices(repo, uuidP, timeP, shortenerP)

	urlServer := handlers.NewUrlServer(urlServices)

	// running services
	urlServer.ListenAndServe(":3000")
}
