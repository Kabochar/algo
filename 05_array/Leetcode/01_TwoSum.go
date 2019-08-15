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
		_, in := m[v]
		if in {
			return []int{m[target-v], i}
		}
		m[target-v] = i
	}

	return []int{0, 0}
}
