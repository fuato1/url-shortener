package qr

import (
	"github.com/fuato1/shorturl/internal/app/qr/commands"
	"github.com/fuato1/shorturl/internal/app/qr/queries"
	"github.com/fuato1/shorturl/internal/domain"
	"github.com/fuato1/shorturl/internal/pkg/qrgen"
	"github.com/fuato1/shorturl/internal/pkg/time"
	"github.com/fuato1/shorturl/internal/pkg/uuid"
)

// QR queries and commands
type QRQueries struct {
	GetAllQRsHandler queries.GetAllQRsRequestHandler
}

type QRCommands struct {
	AddQRHandler commands.AddQRRequestHandler
}

// QR services
type QRServices struct {
	Queries  QRQueries
	Commands QRCommands
}

func NewQRServices(repo domain.QRRepository, uuidP uuid.Provider, timeP time.Provider, qrGenP qrgen.Provider) QRServices {
	return QRServices{
		Queries: QRQueries{
			GetAllQRsHandler: queries.NewGetAllQRsRequestHandler(repo),
		},
		Commands: QRCommands{
			AddQRHandler: commands.NewAddQRRequestHandler(repo, uuidP, timeP, qrGenP),
		},
	}
}
