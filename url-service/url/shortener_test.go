package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	sourceUrl := "https://google.com"
	expectedShortUrl := "Jsz4k57oAX"

	shortUrl, err := Shorten(sourceUrl)
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, shortUrl, expectedShortUrl)
}
