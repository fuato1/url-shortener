package qr

import (
	"time"

	"github.com/boombuler/barcode"
	"github.com/google/uuid"
)

type QR struct {
	Id        uuid.UUID
	Source    string
	Code      *barcode.Barcode
	CreatedAt time.Time
}
