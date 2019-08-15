package String

import (
	"strings"
)

// 翻转字符串
func reverseWords(s string) string {
	// 使用内置函数进行切割
	str := strings.Split(s, " ")
	for i, s := range str {
		// 将切割后的字符串单独翻转
		str[i] = revers(s)
	}

	return strings.Join(str, " ")
}

func revers(s string) string {
	// 为什么要转换为 byte类型？
	// 因为 字符串类型不允许修改呀~~~
	bytes := []byte(s)
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}

	return string(bytes)
}
