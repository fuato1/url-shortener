package api

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/juanjoss/shorturl/cache"
)

type server struct {
	router *mux.Router
	port   string
	fs     http.FileSystem
	cache  cache.Cache
}

func NewServer(port string, embedFS embed.FS) *server {
	// build FS
	buildFS, err := fs.Sub(embedFS, "views")
	if err != nil {
		log.Fatal(err)
	}

	return &server{
		router: mux.NewRouter(),
		port:   port,
		fs:     http.FS(buildFS),
		cache:  cache.New(),
	}
}

func (s *server) ListenAndServe() {
	// register routes
	s.router.HandleFunc("/", http.FileServer(s.fs).ServeHTTP).Methods(http.MethodGet)
	s.router.HandleFunc("/url", s.getAll).Methods(http.MethodGet)
	s.router.HandleFunc("/url", s.shortenURL).Methods(http.MethodPost)

	s.router.HandleFunc("/{id}", s.resolve).Methods(http.MethodGet)
	s.router.HandleFunc("/qr/{id}", s.getQR).Methods(http.MethodGet)

	// create logging and recovery middlewares
	loggedRouter := handlers.LoggingHandler(os.Stdout, s.router)
	recoveryRouter := handlers.RecoveryHandler()(loggedRouter)

	// run the http server
	if s.port == "" {
		log.Printf("no APP_PORT specified, defaulting service port to 8080")
		s.port = ":8080"
	} else {
		log.Printf("shortener service running on port %s", s.port)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.port), recoveryRouter))
}
