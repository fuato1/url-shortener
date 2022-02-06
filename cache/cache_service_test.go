package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCacheService = &CacheService{}

func init() {
	testCacheService = InitCache()
}

func TestCacheInit(t *testing.T) {
	assert.True(t, testCacheService.redisClient != nil)
}

func TestSetAndGet(t *testing.T) {
	originalLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	// save url mapping
	SaveUrl(shortUrl, originalLink, userId)

	// retrieve initial URL
	initialUrl := GetInitialUrl(shortUrl)

	assert.Equal(t, initialUrl, originalLink)
}
