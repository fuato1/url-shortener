package commands

import (
	"log"

	"github.com/fuato1/shorturl/pkg/qrgen"
	"github.com/fuato1/shorturl/pkg/time"
	"github.com/fuato1/shorturl/pkg/uuid"
	"github.com/fuato1/shorturl/qr-service/model"
	"github.com/fuato1/shorturl/qr-service/ports"
)

type AddQrRequest struct {
	Source string
}

type AddQrRequestHandler interface {
	Handle(req AddQrRequest) error
}

type addQrRequestHandler struct {
	uuidProvider  uuid.Provider
	timeProvider  time.Provider
	qrGenProvider qrgen.Provider
	repo          ports.QrRepository
}

func NewAddQrRequestHandler(repo ports.QrRepository, uuidProvider uuid.Provider, timeProvider time.Provider, qrGenProvider qrgen.Provider) *addQrRequestHandler {
	return &addQrRequestHandler{
		repo:          repo,
		uuidProvider:  uuidProvider,
		timeProvider:  timeProvider,
		qrGenProvider: qrGenProvider,
	}
}

func (rh addQrRequestHandler) Handle(req AddQrRequest) error {
	qr := model.QR{
		Id:        rh.uuidProvider.NewUUID(),
		Source:    req.Source,
		Code:      rh.qrGenProvider.Generate(req.Source),
		CreatedAt: rh.timeProvider.Now(),
	}

	err := rh.repo.Add(qr)
	if err != nil {
		log.Fatalf("unable to Add Qr: %v", err)
		return err
	}

	return nil
}
