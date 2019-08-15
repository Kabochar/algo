package Array

// 两数之和等于给定目标数字
func TwoSum(arr []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range arr {
		_, in := tmp[v]
		if in {
			return []int{tmp[target-v], i}
		}
		tmp[target-v] = i
	}

	return []int{0, 0}
}
