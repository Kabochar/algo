package main

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (ln *ListNode) GetNext() *ListNode {
	return ln.next
}

func (ln *ListNode) GetValue() interface{} {
	return ln.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		NewListNode(0), 0,
	}
}

//节点前添加
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newNode := NewListNode(v)
	oldNode := p.next
	p.next = newNode
	newNode.next = oldNode
	l.length++

	return true
}

// 节点后添加
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == l.head {
		return false
	}

	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.length++

	return true
}

// 插入到链表头
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// 插入到链表尾
func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}

	return l.InsertAfter(cur, v)
}

// 根据下标查找
func (l *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= l.length {
		return nil
	}

	cur := l.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}

	return cur
}

// 删除节点
// PS：slow，fast 指针，中间夹杂 目标节点，删除并连接前后即可
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}

	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	l.length--

	return true
}

func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

// 反转
// PS：起码得有两个节点
func (l *LinkedList) Reverse() bool {
	if nil == l.head || nil == l.head.next || nil == l.head.next.next {
		return false
	}

	var pre *ListNode = nil
	cur := l.head.next
	for cur != nil {
		temp := cur.next
		cur.next = pre
		pre = cur
		cur = temp
	}

	l.head.next = pre

	return true
}

// 是否有环
// 使用 slow，fast pos 指针循环查找
func (l *LinkedList) HasCircle() bool {
	if nil != l.head {
		slow := l.head
		fast := l.head

		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}

	return false
}

// 查找中间节点
func (l *LinkedList) FindMiddleNode() *ListNode {
	if nil == l.head || nil == l.head.next {
		return nil
	}
	if nil == l.head.next.next { // 只有三个结点时直接返回中间节点
		return l.head.next
	}

	slow, fast := l.head, l.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

// 删除第 N个节点
func (l *LinkedList) DeleteBottomN(n int) bool {
	if n < 0 || nil == l.head || nil == l.head.next {
		return false
	}

	fast := l.head
	for i := 1; i < n && nil != fast; i++ {
		fast = fast.next
	}

	if nil == fast {
		return false
	}

	slow := l.head
	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next

	return true
}

func InitLinkedList() {
	list := NewLinkedList()

	list.InsertToTail(1)
	list.InsertToTail(2)
	list.Print()

	list.Reverse()
	list.Print()
}

func main() {
	InitLinkedList()
}
