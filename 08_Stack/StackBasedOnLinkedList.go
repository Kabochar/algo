package main

import "fmt"

// 链表实现栈
type Node struct {
	next *Node
	val  interface{}
}

type LinkedListStack struct {
	topNode *Node
}

// 新建
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{
		nil,
	}
}

// 是否为空
func (ls *LinkedListStack) IsEmpty() bool {
	if ls.topNode == nil {
		return true
	}

	return false
}

// 压栈
func (ls *LinkedListStack) Push(v interface{}) {
	ls.topNode = &Node{
		next: ls.topNode,
		val:  v,
	}
}

// 出栈
func (ls *LinkedListStack) Pop() interface{} {
	if ls.IsEmpty() {
		return nil
	}

	val := ls.topNode.val
	ls.topNode = ls.topNode.next

	return val
}

// 栈顶元素
func (ls *LinkedListStack) Top() interface{} {
	if ls.IsEmpty() {
		return nil
	}

	return ls.topNode.val
}

// 清空栈
func (ls *LinkedListStack) Flush() {
	ls.topNode = nil
}

// 打印
func (ls *LinkedListStack) Print() {
	if ls.IsEmpty() {
		fmt.Println("Empty Stack")
	} else {
		cur := ls.topNode
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
