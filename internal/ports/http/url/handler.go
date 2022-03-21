package url

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fuato1/shorturl/internal/app/url"
	"github.com/fuato1/shorturl/internal/app/url/commands"
)

type Handler struct {
	urlServices url.URLServices
}

func NewHandler(urlServices url.URLServices) *Handler {
	return &Handler{urlServices: urlServices}
}

func (h *Handler) GetAll(w http.ResponseWriter, _ *http.Request) {
	urls, err := h.urlServices.Queries.GetAllURLsHandler.Handle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(urls)
	if err != nil {
		return
	}
}

type AddUrlRequestModel struct {
	Source string `json:"source" binding:"required"`
	UserId string `json:"userId" binding:"required"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var url AddUrlRequestModel

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	err = h.urlServices.Commands.AddURLHandler.Handle(commands.AddURLRequest{
		Source: url.Source,
		UserId: url.UserId,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
