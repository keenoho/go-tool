package tool

import (
	"math/rand"
	"sort"
)

// StringInList 字符串数组中是否存在目标字符串
func StringInList(target string, list []string) bool {
	sort.Strings(list)
	index := sort.SearchStrings(list, target)
	return index < len(list) && list[index] == target
}

// 随机字符串
func StringRandom(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+~=-")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
