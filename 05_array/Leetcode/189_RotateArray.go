package Leetcode

// 要求：把数据旋转k步
// 注意：边界处理问题
func rotate(nums []int, k int) {
	lens := len(nums)

	// 截取位置
	if k > lens {
		k = k % lens
	}
	// 如果旋转 k==0，返回；k==len(arr), 返回
	if k == 0 || k == lens {
		return
	}

	reverse(nums, 0, lens-1) // 整体翻转
	reverse(nums, 0, k-1)    // 前 k-1位翻转
	reverse(nums, k, lens-1) // k - lens 位翻转
}

// 翻转
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
