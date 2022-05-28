package queries

import (
	"log"

	"github.com/fuato1/shorturl/qr-service/model"
	"github.com/fuato1/shorturl/qr-service/ports"
)

type GetAllQrsResult struct {
	Qr model.QR
}

type GetAllQrsRequestHandler interface {
	Handle() ([]GetAllQrsResult, error)
}

type getAllQrsRequestHandler struct {
	repo ports.QrRepository
}

func NewGetAllQrsRequestHandler(repo ports.QrRepository) *getAllQrsRequestHandler {
	return &getAllQrsRequestHandler{repo: repo}
}

func (rh getAllQrsRequestHandler) Handle() ([]GetAllQrsResult, error) {
	Qrs, err := rh.repo.GetAll()
	if err != nil {
		log.Fatalf("unable to GetAll Qr codes: %v", err)
		return []GetAllQrsResult{}, err
	}

	var results []GetAllQrsResult
	for _, qr := range Qrs {
		results = append(results, GetAllQrsResult{Qr: qr})
	}

	return results, nil
}
