package main

// 冒泡排序
// a 数组，n 数组长度
func BubbleSort(arr []int, n int) {
	if n < 1 {
		return
	}

	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break // 如果没有数据交换，提前退出
		}
	}
}

// 插入排序
// a 数组，n 数组长度
func InsertionSort(arr []int, n int) {
	if n < 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}

// 选择排序
// a 数组，n 数组长度
func SelectionSort(arr []int, n int) {
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	ab := []int{1, 5, 4, 6, 12, 4, 3}
	// BubbleSort(ab, len(ab))
	// InsertionSort(ab, len(ab))
	SelectionSort(ab, len(ab))
}
