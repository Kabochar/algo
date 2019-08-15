package Array

// 查找 1-100缺失的数字
func FindMissingNum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	target := 0 // 1-100总和
	for i := 1; i <= 100; i++ {
		target += i
	}

	for _, v := range arr {
		target -= v
	}

	return target
}
