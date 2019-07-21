package Leetcode

import (
	"fmt"
)

// MinStack：可以返回最小值的栈
type MinStack struct {
	stack []item
}

// 栈元素
// min 记录当前栈中最小值
// x 记录值
type item struct {
	min int
	x   int
}

// Constructor 构造 MinStack
func Constructor() MinStack {
	return MinStack{}
}

// 压栈，存入数据
func (ms *MinStack) Push(x int) {
	min := x
	if len(ms.stack) > 0 && ms.GetMin() < x {
		min = ms.GetMin()
	}
	ms.stack = append(ms.stack,
		item{
			min: min,
			x:   x,
		})
}

// 出栈，去除最后一个元素
// 疑问：移出栈顶的 min 元素，为什么还能获得 剩下栈 min 值？？
// 答：stack存储是一个 struct 类型，记录着进栈时的 min 值
func (ms *MinStack) Pop() {
	ms.stack = ms.stack[:len(ms.stack)-1]
}

// Top 返回栈中最大值
func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1].x
}

// GetMin 返回最小值
func (ms *MinStack) GetMin() int {
	return ms.stack[len(ms.stack)-1].min
}

func (ms *MinStack) Print() {
	if len(ms.stack) == 0 {
		fmt.Println("Empty Stack")
	} else {
		for _, v := range ms.stack {
			fmt.Print(v.x, "-", v.min, " -> ")
		}
	}
}
