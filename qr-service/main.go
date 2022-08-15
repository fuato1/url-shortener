package main

import (
	"os"

	grpc "github.com/juanjoss/qr-service/handlers/grpc/server"
	"github.com/juanjoss/qr-service/handlers/http"
)

func main() {
	// running service
	done := make(chan bool)

	go func() {
		grpc.ListenAndServe(os.Getenv("GRPC_SERVER_PORT"))
	}()

	go func() {
		http.ListenAndServe(os.Getenv("APP_PORT"))
	}()

	<-done
}
