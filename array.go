package tool

// 字符串数组中是否存在目标字符串
func ArrayContainString(list []string, target string) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}
