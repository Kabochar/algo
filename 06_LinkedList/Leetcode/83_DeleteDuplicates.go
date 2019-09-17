package Leetcode

// 要求：移除重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	// 健壮性判断
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		// 如果值相同，切断Next往后移动
		// 否则，cur移动到下一位
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
