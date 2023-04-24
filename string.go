package tool

import "sort"

// StringInList 字符串数组中是否存在目标字符串
func StringInList(target string, list []string) bool {
	sort.Strings(list)
	index := sort.SearchStrings(list, target)
	return index < len(list) && list[index] == target
}
