package qr

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fuato1/shorturl/internal/app/qr"
	"github.com/gorilla/mux"
)

const qrsPath = "/qr"

type QrServer struct {
	services qr.QrServices
	router   *mux.Router
}

func NewQrServer(qrServices qr.QrServices) *QrServer {
	httpServer := &QrServer{services: qrServices}
	httpServer.router = mux.NewRouter()
	httpServer.registerRoutes()
	http.Handle(qrsPath, httpServer.router)

	return httpServer
}

func (s *QrServer) registerRoutes() {
	// qr queries
	s.router.HandleFunc(qrsPath, NewHandler(s.services).GetAll).Methods("GET")

	// qr commands
	s.router.HandleFunc(qrsPath, NewHandler(s.services).Add).Methods("POST")
}

func (s *QrServer) ListenAndServe(port string) {
	fmt.Println("services running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
