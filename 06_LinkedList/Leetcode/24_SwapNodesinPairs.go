package Leetcode

// 要点：两两交换节点位置
// PS：使用递归，不断推进下去，最后返回新节点
//
// 	src LinkedList: 1-> 2-> 3-> 4
// 	-> 1(head),2(newHead)
// 	->   3(head),4(newHead)
// 	->     4(head)
// 	-------
// 	1: 4->3 return 4...
// 	2: 2->1->4->3 return 2...
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next) // 递归下去
	newHead.Next = head

	return newHead
}
