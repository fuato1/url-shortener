package api

import (
	"encoding/json"
	"net/http"
	neturl "net/url"

	"github.com/juanjoss/shorturl/pkg/shortener"
)

func (s *server) shortenURL(w http.ResponseWriter, r *http.Request) {
	_, span := s.tracer.Start(r.Context(), "shortenURL")
	defer span.End()

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

	s.metrics.shorturlRequestsCounter.Add(r.Context(), 1)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"base_url": "http://localhost/",
		"id":       id,
	})
}
