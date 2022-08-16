package url

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/itchyny/base58-go"
)

func Shorten(source string) (string, error) {
	urlHashBytes, err := toSHA256(source)
	if err != nil {
		return "", err
	}

	bigInt := new(big.Int).SetBytes(urlHashBytes).Uint64()

	finalString, err := toBase58([]byte(fmt.Sprintf("%d", bigInt)))
	if err != nil {
		return "", err
	}

	return finalString[:8], nil
}

func toSHA256(s string) ([]byte, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(s))
	if err != nil {
		return []byte{}, err
	}

	return hash.Sum(nil), nil
}

func toBase58(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}
