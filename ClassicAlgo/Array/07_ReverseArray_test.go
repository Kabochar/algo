package Array

import (
	"testing"
)

func TestReverseArray(t *testing.T) {
	lens := 50
	arr := make([]int, lens)
	for i := 0; i < lens; i++ {
		arr[i] = i + 1
	}
	t.Log("arr = ", arr)

	rlt := ReverseArray(arr)
	t.Log("rlt = ", rlt)
}
