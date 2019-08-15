package Array

func ReverseArray(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		// 指针翻转对应数值
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return arr
}
