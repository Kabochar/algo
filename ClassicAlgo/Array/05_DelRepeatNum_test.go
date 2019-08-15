package Array

import "testing"

func TestDelRepeatNum(t *testing.T) {
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		if i == 3 || i == 4 || i == 5 { // 模拟重复数据
			arr[i] = -5
			continue
		}
		arr[i] = i + 1
	}
	t.Log("arr = ", arr)

	rlt := DelRepeatNum(arr)
	t.Log("rlt = ", rlt)
}
