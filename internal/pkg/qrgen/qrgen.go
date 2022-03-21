package qrgen

import (
	"log"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Provider interface {
	Generate(data string) *barcode.Barcode
}

type qrGenProvider struct{}

func NewQRGenProvider() *qrGenProvider {
	return &qrGenProvider{}
}

func (qg qrGenProvider) Generate(data string) *barcode.Barcode {
	encodedData, err := qr.Encode(data, qr.L, qr.Auto)
	if err != nil {
		log.Fatalf("unable to encode data: %v", err)
	}

	qrCode, err := barcode.Scale(encodedData, 512, 512)
	if err != nil {
		log.Fatalf("unable to scale encoded data: %v", err)
	}

	// png.Encode(w, qr) -> this for the port
	return &qrCode
}
