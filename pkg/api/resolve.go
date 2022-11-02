package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *server) resolve(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println(id)

	url, err := s.cache.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(url)

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
