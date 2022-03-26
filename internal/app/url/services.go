package url

import (
	"github.com/fuato1/shorturl/internal/app/url/commands"
	"github.com/fuato1/shorturl/internal/app/url/queries"
	"github.com/fuato1/shorturl/internal/domain"
	"github.com/fuato1/shorturl/internal/pkg/shortener"
	"github.com/fuato1/shorturl/internal/pkg/time"
	"github.com/fuato1/shorturl/internal/pkg/uuid"
)

// Url queries and commands
type UrlQueries struct {
	GetAllUrlsHandler queries.GetAllUrlsRequestHandler
}

type UrlCommands struct {
	AddUrlHandler commands.AddUrlRequestHandler
}

// Url services
type UrlServices struct {
	Queries  UrlQueries
	Commands UrlCommands
}

func NewUrlServices(repo domain.UrlCacheRepository, uuidP uuid.Provider, timeP time.Provider, shortenerP shortener.Provider) UrlServices {
	return UrlServices{
		Queries: UrlQueries{
			GetAllUrlsHandler: queries.NewGetAllUrlsRequestHandler(repo),
		},
		Commands: UrlCommands{
			AddUrlHandler: commands.NewAddUrlRequestHandler(repo, uuidP, timeP, shortenerP),
		},
	}
}
