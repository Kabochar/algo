package Array

import (
	"testing"
)

func TestFindMissingNum(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		if i == 50 { // 缺失数字
			arr[i] = 0
			continue
		}
		arr[i] = i + 1
	}
	t.Log("arr = ", arr)

	rlt := FindMissingNum(arr)
	t.Log("rlt = ", rlt)
}
