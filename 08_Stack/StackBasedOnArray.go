package main

import "fmt"

// 数组实现栈
type ArrayStack struct {
	data []interface{}
	top  int
}

// 新建栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 20),
		top:  -1,
	}
}

// 是否为空
func (as *ArrayStack) IsEmpty() bool {
	if as.top < 0 {
		return true
	}

	return false
}

// 压栈
func (as *ArrayStack) Push(v interface{}) {
	if as.top < 0 {
		as.top = 0
	} else {
		as.top += 1
	}

	if as.top > len(as.data)-1 {
		as.data = append(as.data, v)
	} else {
		as.data[as.top] = v
	}
}

// 退栈
func (as *ArrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}
	value := as.data[as.top]
	as.top -= 1

	return value
}

// 栈顶元素
func (as *ArrayStack) Top() interface{} {
	if as.IsEmpty() {
		return nil
	}

	return as.data[as.top]
}

// 清空栈
func (as *ArrayStack) Flush() {
	as.top = -1
}

// 打印
func (as *ArrayStack) Print() {
	if as.IsEmpty() {
		fmt.Println("Empty Stack")
	} else {
		for i := as.top; i >= 0; i-- {
			fmt.Println(as.data[i])
		}
	}
}
