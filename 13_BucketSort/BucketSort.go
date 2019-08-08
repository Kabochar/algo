package BucketSort

import (
	Sort "algo/12_Merge-QuickSort"
	"fmt"
)

// 桶排序

// 获取待排序数组中的最大值
func getMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func BucketSort(arr []int) {
	num := len(arr)
	if num <= 1 {
		return
	}
	max := getMax(arr)
	buckets := make([][]int, num) // 二维切片

	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max                // 桶序号
		buckets[index] = append(buckets[index], arr[i]) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			Sort.QuickSort(buckets[i]) // 桶内做快速排序
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}

}

// 桶排序简单实现
// == c 是干嘛的？？？
func BucketSortSimple(arr []int) {
	if len(arr) <= 1 {
		return
	}

	array := make([]int, getMax(arr)+1)
	for i := 0; i < len(arr); i++ {
		array[arr[i]]++
	}
	fmt.Println(array)

	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(arr, c)
}
