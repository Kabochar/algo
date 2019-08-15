package CartoonAlgo

// 判断一个数是否为 2的整数次幂
func IsPowerOf2(num int) bool {
	return (num&num - 1) == 0
}
