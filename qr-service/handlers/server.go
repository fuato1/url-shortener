package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fuato1/shorturl/qr-service/services"
	"github.com/gorilla/mux"
)

const qrsPath = "/qr"

type QrServer struct {
	services services.QrServices
	router   *mux.Router
}

func NewQrServer(qrServices services.QrServices) *QrServer {
	httpServer := &QrServer{services: qrServices}
	httpServer.router = mux.NewRouter()
	httpServer.registerRoutes()
	http.Handle(qrsPath, httpServer.router)

	return httpServer
}

func (s *QrServer) registerRoutes() {
	// qr queries
	s.router.HandleFunc(qrsPath, NewHandler(s.services).GetAll).Methods(http.MethodGet)

	// qr commands
	s.router.HandleFunc(qrsPath, NewHandler(s.services).Add).Methods(http.MethodPost)
}

func (s *QrServer) ListenAndServe(port string) {
	fmt.Println("qr service running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
