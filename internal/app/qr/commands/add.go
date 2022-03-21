package commands

import (
	"log"

	"github.com/fuato1/shorturl/internal/domain"
	"github.com/fuato1/shorturl/internal/domain/qr"
	"github.com/fuato1/shorturl/internal/pkg/qrgen"
	"github.com/fuato1/shorturl/internal/pkg/time"
	"github.com/fuato1/shorturl/internal/pkg/uuid"
)

type AddQRRequest struct {
	Source string
}

type AddQRRequestHandler interface {
	Handle(req AddQRRequest) error
}

type addQRRequestHandler struct {
	uuidProvider  uuid.Provider
	timeProvider  time.Provider
	qrGenProvider qrgen.Provider
	repo          domain.QRRepository
}

func NewAddQRRequestHandler(repo domain.QRRepository, uuidProvider uuid.Provider, timeProvider time.Provider, qrGenProvider qrgen.Provider) *addQRRequestHandler {
	return &addQRRequestHandler{
		repo:          repo,
		uuidProvider:  uuidProvider,
		timeProvider:  timeProvider,
		qrGenProvider: qrGenProvider,
	}
}

func (rh addQRRequestHandler) Handle(req AddQRRequest) error {
	qr := qr.QR{
		Id:        rh.uuidProvider.NewUUID(),
		Source:    req.Source,
		Code:      rh.qrGenProvider.Generate(req.Source),
		CreatedAt: rh.timeProvider.Now(),
	}

	err := rh.repo.Add(qr)
	if err != nil {
		log.Fatalf("unable to Add QR: %v", err)
		return err
	}

	return nil
}
