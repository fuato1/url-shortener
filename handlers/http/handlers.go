package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"net/http"
	neturl "net/url"

	"github.com/gorilla/mux"
	grpc "github.com/juanjoss/shorturl/handlers/grpc/client"
	"github.com/juanjoss/shorturl/repository"
	"github.com/juanjoss/shorturl/url"
)

func GetAll(w http.ResponseWriter, _ *http.Request) {
	urls, err := repository.Get().GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(urls)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err.Error())
		return
	}
}

func Resolve(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	url, err := repository.Get().Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err.Error())
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func GetQR(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	qr, err := grpc.GenerateQR("http://localhost:3000/" + id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	image, _, err := image.Decode(bytes.NewReader(qr))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	png.Encode(w, image)
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request map[string]string

	json.NewDecoder(r.Body).Decode(&request)
	source := request["source"]

	_, err := neturl.ParseRequestURI(source)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	id, err := url.Shorten(source)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = repository.Get().Add(id, source)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"base_url": "http://localhost:3000/",
		"id":       id,
	})
}
