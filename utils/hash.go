package utils

import (
	"github.com/matthewhartstonge/argon2"
)

type Hash struct{}

func (hash *Hash) HashEncoded(value string) (string, error) {
	argon := argon2.DefaultConfig()
	encoded, err := argon.HashEncoded([]byte(value))
	encodedStr := ""
	if err == nil {
		encodedStr = string(encoded)
	}
	return encodedStr, err
}

func (hash *Hash) VerifyEncoded(value string, hashedValue string) (bool, error) {
	return argon2.VerifyEncoded([]byte(value), []byte(hashedValue))
}
