package Array

// 无序数组里面查找最大最小值
func FindMaxAndMin(arr []int) (max int, min int) {
	if len(arr) == 0 {
		return 0, 0
	}
	// 假设 arr[0] 就是 max , min 值
	max = arr[0]
	min = arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return max, min
}
