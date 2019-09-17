package Leetcode

// 暴力破解
// 硬遍历，成功后返回
func twoSum(nums []int, target int) []int {
	lnt := len(nums)
	for i := 0; i < lnt; i++ {
		for j := i; j < lnt; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

// 使用map存储所需数据
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		// 判断 v 是否存在于 map 中
		_, in := m[v]
		// 如果存在，说明有搭配的数据
		if in {
			return []int{m[target-v], i}
		}
		// 如果不存在，key记录 目标数值-当前元素，value 记录下标
		m[target-v] = i
	}

	return []int{0, 0}
}
