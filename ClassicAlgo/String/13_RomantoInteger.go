package String

// 罗马数转整数
// 注意点：
// 1，从右往左遍历
// 2，当出现小数在大数左边，减去对应的小数
// 3，其他，总额加上对应的数值
func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// last 是大数小数的中间人，如果小数在大数左边，执行 if 语句
	last := 0
	// 从右往左遍历，Why？ 解决出现 IV 这样需要减法的操作
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]

		sign := 1
		// 小数在大数左边，需要减去这个小数
		if temp < last {
			sign = -1
		}

		// 如果是正常数值，直接相加
		res += sign * temp

		last = temp
	}

	return res
}
