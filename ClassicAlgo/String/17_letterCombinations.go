package String

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	// 保证有两个字符串
	ret := []string{""}

	// 让digits中所有的数字都有机会进行迭代
	for i := 0; i < len(digits); i++ {
		temp := make([]string, 0)
		// 让ret中的每个元素都能添加新的字母。
		for j := 0; j < len(ret); j++ {
			// 把digits[i]所对应的字母，多次添加到ret[j]的末尾
			for k := 0; k < len(m[digits[i]]); k++ {
				// ret[j]+m[digits[i]][k], 第一次 j 循环时，ret内部没有内容
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}

		ret = temp
	}

	return ret
}
