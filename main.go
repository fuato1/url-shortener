package main

import (
	"embed"
	"os"

	"github.com/juanjoss/shorturl/handlers/http"
)

//go:embed views/*
var fs embed.FS

func main() {
	// running service
	http.ListenAndServe(os.Getenv("APP_PORT"), fs)
}
