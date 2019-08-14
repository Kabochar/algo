package CartoonAlgo

// 去除 K个数字后的最小值
// 怎么能使去除后最小?
// KEY: 把在高位比其右侧一位的数值大的数字删除
// Time: O(n)
// Space: O(n)
// 参考链接: https://blog.csdn.net/MrLiar17/article/details/86686160
func removeKDigits(num string, k int) {
	n := 0
	lens := len(num)
	rlt := make([]rune, lens-k) // 数组存放
	for i := 0; i < lens; i++ {
		cur := num[i]
		for n > 0 && cur < num[n-1] && k > 0 { // 按照要求删除数据
			k-- // 需要删除长度减一
			n-- // 字符串减一
		}
		rlt[n] = rune(cur) // 存放数据
		n++
	}
}
