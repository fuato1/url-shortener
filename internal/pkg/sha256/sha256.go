package sha256

import (
	"crypto/sha256"
	"log"
)

type Provider interface {
	SHA256(s string) ([]byte, error)
}

type sha256Provider struct {
}

func NewSHA256Provider() Provider {
	return sha256Provider{}
}

func (p sha256Provider) SHA256(s string) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		log.Fatalf("unable to write bytes to hash: %v", err)
		return []byte{}, err
	}

	return hash.Sum(nil), nil
}
