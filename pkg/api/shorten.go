package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"os"

	"github.com/juanjoss/shorturl/pkg/shortener"
)

func (s *server) shortenURL(w http.ResponseWriter, r *http.Request) {
	_, span := s.tracer.Start(r.Context(), "shortenURL")
	defer span.End()

	s.metrics.shorturlRequestsCounter.Add(r.Context(), 1)

	var request map[string]string

	json.NewDecoder(r.Body).Decode(&request)
	source := request["source"]

	_, err := neturl.ParseRequestURI(source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := shortener.Shorten(source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.cache.Add(id, source)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"base_url": fmt.Sprintf("http://localhost:%s/", os.Getenv("APP_PORT")),
		"id":       id,
	})
}
