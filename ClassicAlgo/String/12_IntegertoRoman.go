package String

// 参考：https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0012.integer-to-roman/README.md
// 整数转为罗马数。。。玄学
func intToRoman(num int) string {
	// 这里需要自己罗列不同位的可能性
	d := [4][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}

	// 取不同位上的数值，到二位Slice上取罗马数
	return d[3][num/1000] +
		d[2][num/100%10] +
		d[1][num/10%10] +
		d[0][num%10]
}
