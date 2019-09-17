package Leetcode

// 要求：将 nums1 nums2 前 n 项合并，放入 nuns1 中
// 注意：使用 j, k 动态指针进行合并
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 深拷贝 nums1 - 完全拷贝
	temp := make([]int, m)
	copy(temp, nums1)

	// j, k 指针 j->temp, k->nums2
	j, k := 0, 0
	for i := 0; i < len(nums1); i++ {
		// nums2 用完
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		// temp 用完
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		// 比较后，放入nums1中，并移动指针
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}
