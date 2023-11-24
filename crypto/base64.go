package crypto

import "encoding/base64"

func EncryptBase64(input string) string {
	encodedString := base64.StdEncoding.EncodeToString([]byte(input))
	return encodedString
}

func DecryptBase64(input string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return ""
	}
	decodedString := string(decodedBytes)
	return decodedString
}
