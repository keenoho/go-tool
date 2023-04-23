package crypto

import (
	"crypto/rc4"
	"encoding/hex"
)

func EncryptRC4(str string, key string) string {
	plaintext := []byte(str)
	ciphertext := make([]byte, len(plaintext))
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	cipher.XORKeyStream(ciphertext, plaintext)
	return hex.EncodeToString(ciphertext)
}

func DecryptRC4(str string, key string) string {
	decodeStr, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	ciphertext := []byte(decodeStr)
	plaintext := make([]byte, len(ciphertext))
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	cipher.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}
