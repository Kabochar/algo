package Leetcode

//要求：把非零的数前置
//注意：数组前面是非零的
//这里使用了两个指针操作元素
func moveZeroes(nums []int) {
	lens := len(nums)
	i, j := 0, 0 // j->fast, j->last

	for j < lens {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	//此时，i 在 非0数字 的末尾
	//这时候把 i 后面的数值都 置 0 即可
	for i < lens {
		nums[i] = 0
		i++
	}
}