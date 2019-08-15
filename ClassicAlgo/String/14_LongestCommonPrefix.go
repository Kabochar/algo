package String

// 寻找最长相同前缀
func longestCommonPrefix(strs []string) string {
	// 寻找最短模板字符串
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			// 如果出现字符不等，直接截取返回
			// strs[j][i] != byte(r)
			// 给定字符串 j位上的 i位字符与 最短字符串的字符相比较
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

// 寻找最短模板字符串
func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}
