package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func EncryptHMACSHA1(str string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(str))
	res := hex.EncodeToString(mac.Sum(nil))
	return res
}

func EncryptHMACSHA256(str string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	res := hex.EncodeToString(mac.Sum(nil))
	return res
}

func EncryptSHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	res := h.Sum(nil)
	return fmt.Sprintf("%x", res)

}

func EncryptSHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	res := h.Sum(nil)
	return fmt.Sprintf("%x", res)

}
