package url

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fuato1/shorturl/internal/app/url"
	"github.com/gorilla/mux"
)

const urlsPath = "/url"

type UrlServer struct {
	services url.UrlServices
	router   *mux.Router
}

func NewUrlServer(urlServices url.UrlServices) *UrlServer {
	httpServer := &UrlServer{services: urlServices}
	httpServer.router = mux.NewRouter()
	httpServer.registerRoutes()
	http.Handle(urlsPath, httpServer.router)

	return httpServer
}

func (s *UrlServer) registerRoutes() {
	// url queries
	s.router.HandleFunc(urlsPath, NewHandler(s.services).GetAll).Methods("GET")

	// url commands
	s.router.HandleFunc(urlsPath, NewHandler(s.services).Add).Methods("POST")
}

func (s *UrlServer) ListenAndServe(port string) {
	fmt.Println("url service running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
