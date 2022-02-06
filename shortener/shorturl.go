package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(s string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(s))

	return algorithm.Sum(nil)
}

func toBase58(bytes []byte) string {
	encoding := base58.BitcoinEncoding

	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func ShortURL(initialLink string, userId string) string {
	// hashing the original URl
	urlHashBytes := sha256Of(initialLink + userId)

	// creating a big integer
	generatedBigInt := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// converting the big integer to base 58 and taking the first 8 characters
	finalString := toBase58([]byte(fmt.Sprintf("%d", generatedBigInt)))

	return finalString[:8]
}
