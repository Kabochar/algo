package Array

func DelRepeatNum(arr []int) []int {
	// 使用快慢指针
	// 如果数值相同，跳过不处理
	// 如果数值不同，把right的值赋给left的下一位
	left, right := 0, 1
	for ; right < len(arr); right++ {
		if arr[left] == arr[right] {
			continue
		}
		left++
		arr[left] = arr[right]
	}

	return arr[:left+1]
}
