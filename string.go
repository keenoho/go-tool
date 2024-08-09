package tool

import (
	"fmt"
	"math/rand"
)

// 随机字符串
func StringRandom(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+~=-")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 左补齐
func StringPadStart(str string, width int, padStr string) string {
	fillStr := "%" + padStr + "*s"
	return fmt.Sprintf(fillStr, width, str)
}

// 右补齐
func StringPadEnd(str string, width int, padStr string) string {
	mLen := width - len(str)
	return str + StringPadStart("", mLen, padStr)
}
