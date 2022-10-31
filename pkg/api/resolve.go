package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) resolve(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	url, err := s.cache.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
