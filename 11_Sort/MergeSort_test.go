package main

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{3, 2}
	t.Log("排序前：", arr)
	MergeSort(arr)
	t.Log("排序后：", arr)

	arr = []int{5, 4, 3, 2, 1}
	t.Log("排序前：", arr)
	MergeSort(arr)
	t.Log("排序后：", arr)
}
