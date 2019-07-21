package main

import "fmt"

// 栈队列
type ArrayQueue struct {
	queue    []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		make([]interface{}, n),
		n,
		0,
		0,
	}
}

// 入队
func (aq *ArrayQueue) EnQueue(v interface{}) bool {
	if aq.tail == aq.capacity {
		return false
	}

	aq.queue[aq.tail] = v
	aq.tail++

	return true
}

// 出队
func (aq *ArrayQueue) DeQueue() interface{} {
	if aq.head == aq.tail {
		return nil
	}

	val := aq.queue[aq.head]
	aq.head++

	return val
}

// 打印
func (aq *ArrayQueue) Print() string {
	if aq.head == aq.tail {
		return "Empty Queue"
	}
	format := "head"
	for i := aq.head; i < aq.tail; i++ {
		format += fmt.Sprintf(" <- %+v", aq.queue[i])
	}
	format += " <- tail"

	return format
}

// 调试
func InitArrayQueue() {
	aq := NewArrayQueue(5)
	aq.EnQueue(1)
	aq.EnQueue(2)
	fmt.Println(aq.Print())

	aq.DeQueue()
	fmt.Println(aq.Print())

}

func main() {
	InitArrayQueue()
}
