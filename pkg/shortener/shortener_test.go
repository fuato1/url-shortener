package shortener

import (
	"testing"

	"github.com/fuato1/shorturl/pkg/base58"
	"github.com/fuato1/shorturl/pkg/sha256"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	shortenerProvider := NewShortenerProvider(sha256.NewSHA256Provider(), base58.NewBase58Provider())

	sourceUrl := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	expectedShortUrl := "Jsz4k57oAX"

	shortUrl := shortenerProvider.Shorten(sourceUrl, userId)

	assert.Equal(t, shortUrl, expectedShortUrl)
}
