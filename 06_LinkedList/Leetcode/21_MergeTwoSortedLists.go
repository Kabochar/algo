package Leetcode

// 要求：合并两个有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 健壮性判定，一条链为nil，直接返回另一条
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	}

	// 头结点最小的作为 head
	// head 作为 头，cur 作为连接器，不断游走在 l1,l2 之间
	var head, cur *ListNode
	if l1.Val < l2.Val {
		head = l1
		cur = l1
		l1 = l1.Next
	} else {
		head = l2
		cur = l2
		l2 = l2.Next
	}

	// 遍历。终止条件：某链表到末尾
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val { // 值小的连接
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		// 关键，cur 要不断前进
		// 保证 head节点形成完整链表
		cur = cur.Next
	}

	// 连接 l1，l2 最后一个结点
	if nil != l1 {
		cur.Next = l1
	}

	if nil != l2 {
		cur.Next = l2
	}

	return head
}
