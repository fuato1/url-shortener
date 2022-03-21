package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fuato1/shorturl/internal/app"
	"github.com/fuato1/shorturl/internal/ports/http/qr"
	"github.com/fuato1/shorturl/internal/ports/http/url"
	"github.com/gorilla/mux"
)

const urlsPath = "/url"
const qrsPath = "/qr"

type Server struct {
	appServices app.Services
	router      *mux.Router
}

func NewServer(appServices app.Services) *Server {
	httpServer := &Server{appServices: appServices}
	httpServer.router = mux.NewRouter()
	httpServer.registerRoutes()
	http.Handle("/", httpServer.router)

	return httpServer
}

func (s *Server) registerRoutes() {
	// url queries
	s.router.HandleFunc(urlsPath, url.NewHandler(s.appServices.URLServices).GetAll).Methods("GET")

	// url commands
	s.router.HandleFunc(urlsPath, url.NewHandler(s.appServices.URLServices).Add).Methods("POST")

	// qr queries
	s.router.HandleFunc(qrsPath, qr.NewHandler(s.appServices.QRServices).GetAll).Methods("GET")

	// qr commands
	s.router.HandleFunc(qrsPath, qr.NewHandler(s.appServices.QRServices).Add).Methods("POST")
}

func (s *Server) ListenAndServe(port string) {
	fmt.Println("services running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
