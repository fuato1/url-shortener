package main

import (
	"os"

	"github.com/juanjoss/url-service/handlers/http"
)

func main() {
	// running service
	http.ListenAndServe(os.Getenv("APP_PORT"))
}
