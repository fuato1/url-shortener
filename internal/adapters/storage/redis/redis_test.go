package redis

import (
	"testing"

	"github.com/fuato1/shorturl/internal/domain/url"
	"github.com/fuato1/shorturl/internal/pkg/base58"
	"github.com/fuato1/shorturl/internal/pkg/sha256"
	"github.com/fuato1/shorturl/internal/pkg/shortener"
	"github.com/fuato1/shorturl/internal/pkg/time"
	"github.com/fuato1/shorturl/internal/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	redisService      = &Repo{}
	uuidProvider      = uuid.NewUUIDProvider()
	timeProvider      = time.NewTimeProvider()
	shortenerProvider = shortener.NewShortenerProvider(
		sha256.NewSHA256Provider(),
		base58.NewBase58Provider(),
	)
)

func init() {
	redisService = NewRepo()
}

func TestCacheConn(t *testing.T) {
	assert.True(t, redisService.redisClient != nil)
}

func TestSetAndGet(t *testing.T) {
	sourceUrl := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userId := "e0dba740-fc4b-4977-872c-d360239e6b1a"

	url := url.ShortUrl{
		Id:        uuidProvider.NewUUID(),
		Source:    sourceUrl,
		URL:       shortenerProvider.Shorten(sourceUrl, userId),
		CreatedAt: timeProvider.Now(),
	}

	// add url
	redisService.Add(url)

	// get initial url
	initialUrl, err := redisService.Get(url.Source)
	if err != nil {
		t.Fail()
	}

	if initialUrl != url.Source {
		t.Fatalf("urls not match. \ninitialUrl: %v \nexpectedUrl: %v", initialUrl, sourceUrl)
	}
}
