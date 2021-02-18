package main

import "fmt"

/**
25. K 个一组翻转链表
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例：
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5

说明：
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 确定k group
	tail := head
	for i := 0; i < k; i++ {
		if tail == nil {
			return head
		}
		tail = tail.Next // 这里得到的最后一个tail，是下一个k group的头指针
	}
	// 先翻转下一个k group
	// 递归返回的结果，作为pre,刚好和当前翻转链表连接
	pre := reverseKGroup(tail, k)
	// 反转当前k区间
	for i := 0; i < k; i++ {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

func main()  {
	list1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmtList(list1)
	fmt.Println("######")
	res := reverseKGroup(list1, 2)
	fmt.Println("----------")
	fmtList(res)
}