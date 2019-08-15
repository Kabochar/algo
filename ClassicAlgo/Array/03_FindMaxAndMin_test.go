package Array

import "testing"

func TestFindMaxAndMin(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		if i == 25 {
			arr[i] = -5
			continue
		}
		arr[i] = i + 1
	}
	t.Log("arr = ", arr)

	max, min := FindMaxAndMin(arr)
	t.Logf("max = %v, min = %v", max, min)
}
