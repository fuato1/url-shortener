package api

import (
	"encoding/json"
	"net/http"
)

func (s *server) getAll(w http.ResponseWriter, r *http.Request) {
	urls, err := s.cache.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
