package ports

import (
	"github.com/fuato1/shorturl/internal/app"
	"github.com/fuato1/shorturl/internal/ports/http"
)

type Services struct {
	Server *http.Server
}

func NewServices(appServices app.Services) Services {
	return Services{Server: http.NewServer(appServices)}
}
