package url

import (
	"time"

	"github.com/google/uuid"
)

type ShortUrl struct {
	Id        uuid.UUID
	Source    string
	URL       string
	CreatedAt time.Time
}
