package String

// 翻转 k个字符
func reverseStr(s string, k int) string {
	bstr := []byte(s)
	// 注意，题目要求是每隔2个字符翻转一次
	for i := 0; i < len(s); i += 2 * k {
		// 找到需要翻转的右边界
		j := min(i+k, len(s))
		// 翻转
		reverse(bstr[i:j])
	}

	return string(bstr)
}

func reverse(bytes []byte) {
	// 老样子，快慢指针走起
	left, right := 0, len(bytes)-1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



