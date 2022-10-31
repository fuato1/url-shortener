package main

import (
	"embed"
	"os"

	"github.com/juanjoss/shorturl/pkg/api"
)

//go:embed views/*
var fs embed.FS

func main() {
	// creating and running service
	server := api.NewServer(os.Getenv("APP_PORT"), fs)
	server.ListenAndServe()
}
