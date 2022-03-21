package url

import (
	"github.com/fuato1/shorturl/internal/app/url/commands"
	"github.com/fuato1/shorturl/internal/app/url/queries"
	"github.com/fuato1/shorturl/internal/domain"
	"github.com/fuato1/shorturl/internal/pkg/shortener"
	"github.com/fuato1/shorturl/internal/pkg/time"
	"github.com/fuato1/shorturl/internal/pkg/uuid"
)

// URL queries and commands
type URLQueries struct {
	GetAllURLsHandler queries.GetAllURLsRequestHandler
}

type URLCommands struct {
	AddURLHandler commands.AddURLRequestHandler
}

// URL services
type URLServices struct {
	Queries  URLQueries
	Commands URLCommands
}

func NewURLServices(repo domain.URLRepository, uuidP uuid.Provider, timeP time.Provider, shortenerP shortener.Provider) URLServices {
	return URLServices{
		Queries: URLQueries{
			GetAllURLsHandler: queries.NewGetAllURLsRequestHandler(repo),
		},
		Commands: URLCommands{
			AddURLHandler: commands.NewAddURLRequestHandler(repo, uuidP, timeP, shortenerP),
		},
	}
}
