package main

import (
	"fmt"
)

// 链式队列
type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{
		nil,
		nil,
		0,
	}
}

// 入队
func (lq *LinkedListQueue) EnQueue(v interface{}) bool {
	node := &ListNode{v, nil}
	if lq.tail == nil {
		lq.tail = node
		lq.head = node
	} else {
		lq.tail.next = node
		lq.tail = node
	}
	lq.length++

	return true
}

// 出队
func (lq *LinkedListQueue) DeQueue() interface{} {
	if lq.head == nil {
		return nil
	}
	val := lq.head.val
	lq.head = lq.head.next

	return val
}

// 打印
func (lq *LinkedListQueue) Print() string {
	if lq.head == nil {
		return "Empty Queue"
	}

	format := "head "
	for cur := lq.head; cur != nil; cur = cur.next {
		format += fmt.Sprintf("<- %+v ", cur.val)
	}
	format += "<- tail"

	return format
}

// 调试
func InitLinkedListQueue() {
	lq := NewLinkedListQueue()
	lq.EnQueue(4)
	lq.EnQueue(5)
	fmt.Println(lq.Print())

	lq.DeQueue()
	fmt.Println(lq.Print())
}

func main() {
	InitLinkedListQueue()
}
