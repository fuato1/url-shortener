package handler

import (
	"fmt"
	"net/http"

	"github.com/fuato1/shorturl/cache"
	"github.com/fuato1/shorturl/shortener"
	"github.com/labstack/echo"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

type Response struct {
	Msg      string `json:"msg" binding:"required"`
	ShortUrl string `json:"short_url" binding:"required"`
}

func CreateShortURL(c echo.Context) error {
	// validating request and getting parameters
	var creationReq UrlCreationRequest
	if err := c.Bind(&creationReq); err != nil {
		fmt.Sprintf("Error shorting URL. Error: %v\n", err)
	}
	fmt.Printf("new request incoming: \n\n%+v\n\n", creationReq)

	// shorting the URL
	shortUrl := shortener.ShortURL(creationReq.LongUrl, creationReq.UserId)

	// saving URL mapping to cache
	cache.SaveUrl(shortUrl, creationReq.LongUrl, creationReq.UserId)

	// defining response
	host := "http://localhost:3000/"
	res := &Response{
		Msg:      "short URL created succesfully",
		ShortUrl: host + shortUrl,
	}
	fmt.Printf("shortened URL: %+v\n\n", shortUrl)

	return c.JSON(http.StatusOK, res)
}

func HandleRedirect(c echo.Context) error {
	shortUrl := c.Param("shortUrl")
	initialUrl := cache.GetInitialUrl(shortUrl)
	return c.Redirect(302, initialUrl)
}
