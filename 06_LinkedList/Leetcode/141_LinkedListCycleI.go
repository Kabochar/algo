package Leetcode

// 要求：判断链表是否存在环
// 方式一，使用 slow，fast pos 指针
func hasCycle_2(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow == fast
}

// 方式二，使用map
func hasCycle_1(head *ListNode) bool {
	tmap := map[*ListNode]bool{}

	cur := head
	for nil != cur {
		if _, ok := tmap[cur]; ok {
			return true
		} else {
			tmap[cur] = true
		}
		cur = cur.Next
	}

	return false
}
