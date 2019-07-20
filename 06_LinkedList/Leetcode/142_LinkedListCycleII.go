package Leetcode

// 要求：求出入环结点
// 注意：公式->
//		首次相遇点 到 入环点 == 首结点 到 入环点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow != fast { // 如果不存在环，结束
		return nil
	}

	for slow != head { // 首结点，入环点 共同走一步，相遇就是入环点
		slow, head = slow.Next, head.Next
	}

	return slow
}
