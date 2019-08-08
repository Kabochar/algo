package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	t.Log("排序前：", arr)
	BubbleSort(arr, len(arr))
	t.Log("排序后：", arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	t.Log("排序前：", arr)
	InsertionSort(arr, len(arr))
	t.Log("排序后：", arr)
}

func TestSelectionSort(t *testing.T) {
	arr := []int{1, 5, 9, 6, 3, 7, 5, 10}
	t.Log("排序前：", arr)
	SelectionSort(arr, len(arr))
	t.Log("排序后：", arr)
}
