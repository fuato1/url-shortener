package shortener

import (
	"fmt"
	"log"
	"math/big"

	"github.com/fuato1/shorturl/internal/pkg/base58"
	"github.com/fuato1/shorturl/internal/pkg/sha256"
)

type Provider interface {
	Shorten(url string, userId string) string
}

type shortenerProvider struct {
	sha256Provider sha256.Provider
	base58Provider base58.Provider
}

func NewShortenerProvider(sha256Provider sha256.Provider, base58Provider base58.Provider) shortenerProvider {
	return shortenerProvider{sha256Provider: sha256Provider, base58Provider: base58Provider}
}

func (hp shortenerProvider) Shorten(url string, userId string) string {
	// hashing the original URl
	urlHashBytes, err := hp.sha256Provider.SHA256(url + userId)
	if err != nil {
		log.Fatalf("unable to hash original URL: %v", err)
		return ""
	}

	// creating a big integer
	bigInt := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// converting the big integer to base 58 and taking the first 8 characters
	finalString, err := hp.base58Provider.ToBase58([]byte(fmt.Sprintf("%d", bigInt)))
	if err != nil {
		log.Fatalf("unable to convert big integer to base58: %v", err)
		return ""
	}

	return finalString[:8]
}
