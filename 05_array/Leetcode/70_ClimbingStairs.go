package Leetcode

// 要求，一次 1 步，一次 2 步
// 本质：斐波那契数列 1, 1, 2, 3, 5, 8...
// 注意：本数 = 前数 + 前前数
// 从 2 开始执行
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}

	res := make([]int, n+1)
	res[0], res[1] = 1, 1

	for i := 2; i <= n; i++ {
		res[i] = res[i-1] + res[i-2]
	}

	return res[n]
}
