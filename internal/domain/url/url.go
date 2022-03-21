package url

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	Id        uuid.UUID
	Source    string
	URL       string
	CreatedAt time.Time
}
