package main

import (
	"fmt"

	"github.com/fuato1/shorturl/internal/adapters"
	"github.com/fuato1/shorturl/internal/app"
	"github.com/fuato1/shorturl/internal/app/qr"
	"github.com/fuato1/shorturl/internal/app/url"
	"github.com/fuato1/shorturl/internal/pkg/base58"
	"github.com/fuato1/shorturl/internal/pkg/loadbalancer"
	"github.com/fuato1/shorturl/internal/pkg/qrgen"
	"github.com/fuato1/shorturl/internal/pkg/sha256"
	"github.com/fuato1/shorturl/internal/pkg/shortener"
	"github.com/fuato1/shorturl/internal/pkg/time"
	"github.com/fuato1/shorturl/internal/pkg/uuid"
	"github.com/fuato1/shorturl/internal/ports"
)

func main() {
	// running services
	done := make(chan bool)

	// running load balancer
	lb := loadbalancer.NewLoadBalancingProvider()

	// running passive healthcheck routine
	// go lb.HealthCheck()

	runUrlServices()
	runQrServices()

	lb.ListenAndServe()

	<-done
}

func runUrlServices() {
	// creating adapter services
	adapterServices := adapters.NewUrlServices()

	// creating providers
	uuidP := uuid.NewUUIDProvider()
	timeP := time.NewTimeProvider()
	shortenerP := shortener.NewShortenerProvider(
		sha256.NewSHA256Provider(),
		base58.NewBase58Provider(),
	)

	// injecting dependencies
	urlServices := url.NewUrlServices(adapterServices.UrlCacheRepo, uuidP, timeP, shortenerP)

	appServices := app.NewUrlServices(urlServices)

	portServices := ports.NewUrlServices(appServices)

	// running services
	for i := 1; i <= 3; i++ {
		go func(id int) {
			portServices.UrlServer.ListenAndServe(fmt.Sprintf(":300%d", id))
		}(i)
	}
}

func runQrServices() {
	// creating adapter services
	adapterServices := adapters.NewQrServices()

	// creating providers
	uuidP := uuid.NewUUIDProvider()
	timeP := time.NewTimeProvider()
	qrgenP := qrgen.NewQrGenProvider()

	// injecting dependencies
	qrServices := qr.NewQrServices(adapterServices.QrSqlRepo, uuidP, timeP, qrgenP)

	appServices := app.NewQrServices(qrServices)

	portServices := ports.NewQrServices(appServices)

	// running services
	for i := 1; i <= 2; i++ {
		go func(id int) {
			portServices.QrServer.ListenAndServe(fmt.Sprintf(":400%d", id))
		}(i)
	}
}
