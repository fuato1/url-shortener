package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router

func ListenAndServe(port string) {
	router = mux.NewRouter()

	router.HandleFunc("/", http.FileServer(http.Dir("./views")).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/qr", GenerateQR).Methods(http.MethodPost)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Printf("qr service running on port %s", port)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%s", port),
			handlers.RecoveryHandler()(loggedRouter),
		),
	)
}
