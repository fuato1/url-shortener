package services

import (
	"github.com/fuato1/shorturl/pkg/qrgen"
	"github.com/fuato1/shorturl/pkg/time"
	"github.com/fuato1/shorturl/pkg/uuid"
	"github.com/fuato1/shorturl/qr-service/ports"
	"github.com/fuato1/shorturl/qr-service/services/commands"
	"github.com/fuato1/shorturl/qr-service/services/queries"
)

// QR queries and commands
type QrQueries struct {
	GetAllQrsHandler queries.GetAllQrsRequestHandler
}

type QrCommands struct {
	AddQrHandler commands.AddQrRequestHandler
}

type QrServices struct {
	Queries  QrQueries
	Commands QrCommands
}

func NewQrServices(repo ports.QrRepository, uuidP uuid.Provider, timeP time.Provider, qrGenP qrgen.Provider) QrServices {
	return QrServices{
		Queries: QrQueries{
			GetAllQrsHandler: queries.NewGetAllQrsRequestHandler(repo),
		},
		Commands: QrCommands{
			AddQrHandler: commands.NewAddQrRequestHandler(repo, uuidP, timeP, qrGenP),
		},
	}
}
