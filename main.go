package main

import (
	"net/http"

	"github.com/fuato1/shorturl/cache"
	"github.com/fuato1/shorturl/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "URL Shortener")
	})

	e.POST("/createShortUrl", handler.CreateShortURL)

	e.GET("/:shortUrl", handler.HandleRedirect)

	cache.InitCache()

	e.Logger.Fatal(e.Start(":3000"))
}
