package Leetcode

import (
	"container/list"
)

// == 跳过这题，使用到了内置库
type MyQueue struct {
	list *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		list: list.New(),
	}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.list.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	front := q.list.Front()
	res := front.Value.(int)
	q.list.Remove(front)

	return res
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	front := q.list.Front()
	res := front.Value.(int)

	return res
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.list.Len() == 0
}
