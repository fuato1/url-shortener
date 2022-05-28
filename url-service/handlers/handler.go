package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fuato1/shorturl/url-service/services"
	"github.com/fuato1/shorturl/url-service/services/commands"
)

type Handler struct {
	urlServices services.UrlServices
}

func NewHandler(urlServices services.UrlServices) *Handler {
	return &Handler{urlServices: urlServices}
}

func (h *Handler) GetAll(w http.ResponseWriter, _ *http.Request) {
	result, err := h.urlServices.Queries.GetAllUrlsHandler.Handle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(result.Urls)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err.Error())
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

	err = h.urlServices.Commands.AddUrlHandler.Handle(commands.AddUrlRequest{
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
