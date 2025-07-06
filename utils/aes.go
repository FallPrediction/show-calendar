package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"os"
)

type Aes struct {
	key []byte
}

func (a *Aes) Decrypt(data string) string {
	tokenBytes, err := hex.DecodeString(data)
	if err != nil {
		panic("Hex decode token failed: " + err.Error())
	}
	key := []byte(os.Getenv("AES_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := tokenBytes[:nonceSize], tokenBytes[nonceSize:]
	result, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(result)
}

func NewAes() Aes {
	key := os.Getenv("AES_KEY")
	return Aes{[]byte(key)}
}
