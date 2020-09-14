package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"unicode"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	randomBytes := make([]byte, n)
	_, err := rand.Read(randomBytes)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", err
	}

	for i, b := range randomBytes {
		randomBytes[i] = alphabet[b%byte(len(alphabet))]
	}
	return string(randomBytes), nil
}

// generateRandomString generates a random string with an predefined length
func generateRandomString(length int) (string, error) {
	var result string
	for len(result) < length {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		if unicode.IsLetter(rune(n)) {
			result += string(n)
		}
	}
	return result, nil
}

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Println(GenerateRandomString(4))
	}
}
