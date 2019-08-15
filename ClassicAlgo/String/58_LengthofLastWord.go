package String

// 末尾字符串的长度
func lengthOfLastWord(s string) int {
	lens := len(s)
	if lens == 0 {
		return 0
	}

	// 结果计数器
	res := 0
	// 从右往左遍历
	for i := lens - 1; i >= 0; i-- {
		// 如果碰到空格字符，判断后返回
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}
		res++
	}

	return res
}
