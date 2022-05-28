package services

import (
	"github.com/fuato1/shorturl/pkg/shortener"
	"github.com/fuato1/shorturl/pkg/time"
	"github.com/fuato1/shorturl/pkg/uuid"
	"github.com/fuato1/shorturl/url-service/ports"
	"github.com/fuato1/shorturl/url-service/services/commands"
	"github.com/fuato1/shorturl/url-service/services/queries"
)

// Url queries and commands
type UrlQueries struct {
	GetAllUrlsHandler queries.GetAllUrlsRequestHandler
}

type UrlCommands struct {
	AddUrlHandler commands.AddUrlRequestHandler
}

type UrlServices struct {
	Queries  UrlQueries
	Commands UrlCommands
}

func NewUrlServices(repo ports.UrlRepository, uuidP uuid.Provider, timeP time.Provider, shortenerP shortener.Provider) UrlServices {
	return UrlServices{
		Queries: UrlQueries{
			GetAllUrlsHandler: queries.NewGetAllUrlsRequestHandler(repo),
		},
		Commands: UrlCommands{
			AddUrlHandler: commands.NewAddUrlRequestHandler(repo, uuidP, timeP, shortenerP),
		},
	}
}
