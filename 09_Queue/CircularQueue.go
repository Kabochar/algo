package main

import (
	"fmt"
)

// 环型队列
type CircularQueue struct {
	queue    []interface{}
	capacity int
	head     int
	tail     int
}

// 新建环型队列
func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}

	return &CircularQueue{
		make([]interface{}, n),
		n,
		0,
		0,
	}
}

// 是否为空
func (cq *CircularQueue) IsEmpty() bool {
	if cq.head == cq.tail {
		return true
	}

	return false
}

// 是否已满
func (cq *CircularQueue) IsFull() bool {
	if cq.head == (cq.tail+1)%cq.capacity {
		return true
	}

	return false
}

// 进队
func (cq *CircularQueue) EnQueue(v interface{}) bool {
	if cq.IsFull() {
		return false
	}

	cq.queue[cq.tail] = v
	cq.tail = (cq.tail + 1) % cq.capacity

	return true
}

// 出队
func (cq *CircularQueue) DeQueue() interface{} {
	if cq.IsEmpty() {
		return nil
	}

	val := cq.queue[cq.head]
	cq.head = (cq.tail + 1) % cq.capacity

	return val
}

// 打印
func (cq *CircularQueue) Print() string {
	if cq.IsEmpty() {
		return "Empty queue"
	}
	result := "head"
	var i = cq.head
	for true {
		result += fmt.Sprintf(" <- %+v", cq.queue[i])
		i = (i + 1) % cq.capacity
		if i == cq.tail {
			break
		}
	}
	result += " <- tail"

	return result
}

func InitCircularQueue() {
	cq := NewCircularQueue(3)
	cq.EnQueue(2)
	fmt.Println("cq.IsFull(): ", cq.IsFull())
	fmt.Println("cq.IsEmpty(): ", cq.IsEmpty())
	fmt.Println(cq.Print())
}

func main() {
	InitCircularQueue()
}
