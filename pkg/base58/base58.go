package base58

import (
	"log"

	"github.com/itchyny/base58-go"
)

type Provider interface {
	ToBase58(bytes []byte) (string, error)
}

type base58Provider struct {
}

func NewBase58Provider() base58Provider {
	return base58Provider{}
}

func (p base58Provider) ToBase58(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Fatalf("unable to encode bytes: %v", err)
		return "", err
	}

	return string(encoded), nil
}
