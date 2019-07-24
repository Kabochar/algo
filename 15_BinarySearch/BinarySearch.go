package main

// 二分查找
func BinarySearch(arr []int, val int) int {
	lens := len(arr)
	if lens == 0 {
		return -1
	}

	low := 0
	high := lens - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// 递归二分查找
func BinarySearchRecursive(a []int, val int) int {
	lens := len(a)
	if lens == 0 {
		return -1
	}

	return bs(a, val, 0, lens-1)
}

func bs(arr []int, val int, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] == val {
		return mid
	} else if arr[mid] > val {
		return bs(arr, val, low, mid-1)
	} else {
		return bs(arr, val, mid+1, high)
	}
}

// 二分查找变体

// 查找第一个等于给定值的元素
func BinarySearchFirst(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != val {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// 查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == n-1 || arr[mid+1] != val {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearchFirstGT(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (high + low) >> 1
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid != n-1 && arr[mid+1] > val {
				return mid + 1
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

// 查找最后一个小于等于给定值的元素
func BinarySearchLastLT(arr []int, val int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) >> 1
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < val {
				return mid - 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
