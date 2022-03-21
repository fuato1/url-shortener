package queries

import (
	"log"

	"github.com/fuato1/shorturl/internal/domain"
	"github.com/fuato1/shorturl/internal/domain/qr"
)

type GetAllQRsResult struct {
	QR qr.QR
}

type GetAllQRsRequestHandler interface {
	Handle() ([]GetAllQRsResult, error)
}

type getAllQRsRequestHandler struct {
	repo domain.QRRepository
}

func NewGetAllQRsRequestHandler(repo domain.QRRepository) *getAllQRsRequestHandler {
	return &getAllQRsRequestHandler{repo: repo}
}

func (rh getAllQRsRequestHandler) Handle() ([]GetAllQRsResult, error) {
	QRs, err := rh.repo.GetAll()
	if err != nil {
		log.Fatalf("unable to GetAll QR codes: %v", err)
		return []GetAllQRsResult{}, err
	}

	var results []GetAllQRsResult
	for _, qr := range QRs {
		results = append(results, GetAllQRsResult{QR: qr})
	}

	return results, nil
}
