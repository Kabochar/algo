package Array

import (
	"testing"
)

func TestRepeatNum(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		// 模拟重复数字
		if i == 25 || i == 30 {
			arr[i] = 132
			continue
		}
		if i == 26 || i == 31 {
			arr[i] = 111
			continue
		}
		arr[i] = i + 1
	}
	t.Log("arr = ", arr)

	rlt := RepeatNum(arr)
	t.Log(rlt)
}
