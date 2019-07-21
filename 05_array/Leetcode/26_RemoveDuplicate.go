package Leetcode

//要求：O(1)操作，原地修改数组
//@numbers 记录不重复元素的下标
//如果元素不重复，更新到 nums 的 numbers 下标位置
func removeDuplicates(nums []int) int {
	numbers := 0
	for i := range nums {
		if nums[i] != nums[numbers] {
			numbers++
			nums[numbers] = nums[i]
		}
	}

	return numbers + 1
}
