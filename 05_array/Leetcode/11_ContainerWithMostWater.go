package Leetcode

//要求：面积最大
//注意：头尾一起找，高度是按照低的计算
//最好不断移动，寻找最大的可能性
func maxArea(height []int) int {
	//头尾一起找
	i, j := 0, len(height)-1
	max := 0 // 记录 max area

	for i < j { // 终止条件，右边 > 左边
		ha, hb := height[i], height[j]
		h := min(ha, hb) // 低的部分作为高

		// 计算当前位置面积
		area := h * (j - i)
		if area > max {
			max = area
		}

		// 不断移动，寻找最大area
		if ha < hb {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(ha, hb int) int {
	if ha < hb {
		return ha
	} else {
		return hb
	}
}
