package MergeSort

// 排序一个数组
// 先把数组从中间分成前后两部分
// 然后对前后两部分分别排序
// 再将排好序的两部分合并在一起，这样整个数据就都有序了
func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	mergeSort(arr, 0, arrLen-1)
}

func mergeSort(arr []int, start, end int) {
	// 递归终止条件
	if start >= end {
		return
	}
	// 取 start，end 的中间位置 mid
	mid := (start + end) / 2
	// 分支递归
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	// 合并两个数组
	merge(arr, start, mid, end)
}

// 稳定性主要查看 merge 函数，归并是稳定的排序算法
func merge(arr []int, start, mid, end int) {
	tmpArr := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] < arr[j] {
			tmpArr[k] = arr[i]
			i++
		} else {
			tmpArr[k] = arr[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmpArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tmpArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tmpArr)
}
