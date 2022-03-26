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
type QrQueries struct {
	GetAllQrsHandler queries.GetAllQrsRequestHandler
}

type QrCommands struct {
	AddQrHandler commands.AddQrRequestHandler
}

// QR services
type QrServices struct {
	Queries  QrQueries
	Commands QrCommands
}

func NewQrServices(repo domain.QrRepository, uuidP uuid.Provider, timeP time.Provider, qrGenP qrgen.Provider) QrServices {
	return QrServices{
		Queries: QrQueries{
			GetAllQrsHandler: queries.NewGetAllQrsRequestHandler(repo),
		},
		Commands: QrCommands{
			AddQrHandler: commands.NewAddQrRequestHandler(repo, uuidP, timeP, qrGenP),
		},
	}
}
