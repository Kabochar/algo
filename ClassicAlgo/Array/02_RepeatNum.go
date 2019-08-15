package Array

// 查找数组重复数字
func RepeatNum(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// 结果集
	var rlt []int
	// 记录当前数字是否存在
	tmp := make(map[int]bool)
	for _, v := range arr {
		// 数字存在，记录到结果集里
		if tmp[v] == true {
			rlt = append(rlt, v)
			continue
		}

		tmp[v] = true
	}

	return rlt
}
