package String

// 查询给定字符串出现的位置
// 卡片比较，滑动窗口
func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen-nlen; i++ {
		// 卡片比较，如果两张卡片相同，返回当前下标
		// 注意：haystack[i:i+nlen]，截取 [i ~ i+nlen] 长度的卡片内容
		if haystack[i:i+nlen] == needle {
			return i
		}
	}

	return -1
}
