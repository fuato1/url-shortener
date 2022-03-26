package main

import (
	"os"

	"github.com/fuato1/shorturl/internal/adapters"
	"github.com/fuato1/shorturl/internal/app"
	"github.com/fuato1/shorturl/internal/app/qr"
	"github.com/fuato1/shorturl/internal/app/url"
	"github.com/fuato1/shorturl/internal/pkg/base58"
	"github.com/fuato1/shorturl/internal/pkg/qrgen"
	"github.com/fuato1/shorturl/internal/pkg/sha256"
	"github.com/fuato1/shorturl/internal/pkg/shortener"
	"github.com/fuato1/shorturl/internal/pkg/time"
	"github.com/fuato1/shorturl/internal/pkg/uuid"
	"github.com/fuato1/shorturl/internal/ports"
)

func main() {
	// creating adapter services
	adapterServices := adapters.NewServices()

	// creating providers
	uuidP := uuid.NewUUIDProvider()
	timeP := time.NewTimeProvider()
	shortenerP := shortener.NewShortenerProvider(
		sha256.NewSHA256Provider(),
		base58.NewBase58Provider(),
	)
	qrgenP := qrgen.NewQrGenProvider()

	// creating url and qr services
	urlServices := url.NewUrlServices(adapterServices.UrlCacheRepo, uuidP, timeP, shortenerP)
	qrServices := qr.NewQrServices(adapterServices.QrSqlRepo, uuidP, timeP, qrgenP)

	// injecting url and qr services into app
	appServices := app.NewServices(urlServices, qrServices)

	// injecting the app services into the ports
	portServices := ports.NewServices(appServices)

	// running services
	done := make(chan bool)

	go func() {
		portServices.UrlServer.ListenAndServe(":" + os.Getenv("URL_SERVICE_PORT"))
	}()

	go func() {
		portServices.QrServer.ListenAndServe(":" + os.Getenv("QR_SERVICE_PORT"))
	}()

	<-done
}
