package Leetcode

// 要求：处理好进位问题，每个位置上的数值只能 < 10
// 特殊点：末尾 +1，首位 +1
func plusOne(digits []int) []int {
	lens := len(digits)

	// 末尾直接添加
	digits[lens-1]++

	// 末尾+1 进位
	for i := lens - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10 // 进位清0
		digits[i-1]++
	}

	// 首尾+1 进位
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
