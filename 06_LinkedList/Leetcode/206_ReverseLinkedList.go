package Leetcode

// 要求：翻转链表
// 注意：结点之间的连接，转换
// PS：
// 	重要的三个节点：
// 	1，待翻转结点
// 	2，待翻转结点的前驱结点 prev
// 	3，待翻转结点的后驱结点 next
// 解析:
// 分为两部分
// 已翻转区 | 待翻转区
// nil   | 1(head)->2->nil
// 第一次：
// 1(head)->nil(pre) | 2->nil
// 1(pre)->nil | 2(head)->nil
// 第二次操作：
// 2(head)->1(pre)->nil | nil
// 2(pre)->1->nil | nil
// -----完成-----
func reverseList(head *ListNode) *ListNode {
	// 健壮性判断，少于两个结点直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// pre 是所有已经逆转结点的 head
	var pre *ListNode
	for head != nil {
		// 让 temp 指向 head.Next，免得 head.Next 丢失
		temp := head.Next
		// head 成为已经逆转结点的新 head
		head.Next = pre
		// 让 pre 重新成为已经逆转结点的 head
		pre = head
		// 让 head 指向下一个被逆转的结点
		head = temp
	}

	return pre
}
