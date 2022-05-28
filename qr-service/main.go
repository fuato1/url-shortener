package main

import (
	"github.com/fuato1/shorturl/pkg/qrgen"
	"github.com/fuato1/shorturl/pkg/time"
	"github.com/fuato1/shorturl/pkg/uuid"
	"github.com/fuato1/shorturl/qr-service/handlers"
	"github.com/fuato1/shorturl/qr-service/repository"
	"github.com/fuato1/shorturl/qr-service/services"
)

func main() {
	// running services
	done := make(chan bool)

	runQrServices()

	<-done
}

func runQrServices() {
	// creating adapter services
	repo := repository.NewQrRepository()

	// creating providers
	uuidP := uuid.NewUUIDProvider()
	timeP := time.NewTimeProvider()
	qrgenP := qrgen.NewQrGenProvider()

	// injecting dependencies
	qrServices := services.NewQrServices(repo, uuidP, timeP, qrgenP)

	qrServer := handlers.NewQrServer(qrServices)

	// running services
	go qrServer.ListenAndServe(":4000")
}
