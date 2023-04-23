package crypto

import (
	"crypto/rc4"
	"encoding/hex"
	"strings"
)

func EncryptRC4(str string, key string) string {
	plaintext := []byte(str)
	ciphertext := make([]byte, len(plaintext))
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	cipher.XORKeyStream(ciphertext, plaintext)
	return strings.ToUpper(hex.EncodeToString(ciphertext))
}

func DecryptRC4(str string, key string) string {
	plaintext, err := hex.DecodeString(strings.ToUpper(str))
	if err != nil {
		panic(err)
	}
	var decrypted []byte
	ciphertext := make([]byte, len(plaintext))
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	decrypted = make([]byte, len(ciphertext))
	cipher.XORKeyStream(decrypted, ciphertext)
	return string(decrypted)
}
