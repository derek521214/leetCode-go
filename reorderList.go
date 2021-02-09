package main

import "fmt"
/**
143. 重排链表
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head == nil || head.Next==nil {
		return
	}
	fast , slow := head, head
	var temp *ListNode
	depth := 0
	for fast != nil && fast.Next != nil {
		temp = slow
		slow = slow.Next
		fast = fast.Next.Next
		depth += 2
	}
	if fast != nil {
		depth++
	}
	var mid *ListNode
	if  depth % 2 == 1 {
		mid = temp.Next.Next
		temp.Next.Next = nil
	} else {
		mid = temp.Next
		temp.Next = nil
	}
	mid = reversalList(mid)
	l1 , l2 := head, mid
	var l1tmp, l2tmp *ListNode
	 for l1 != nil && l2 != nil {
		l1tmp = l1.Next
		l2tmp = l2.Next

		l1.Next = l2
		l1 = l1tmp

		l2.Next = l1
		l2 = l2tmp
	 }
}

func reversalList(node *ListNode) *ListNode {
	res := node
	var pre *ListNode
	for res != nil {
		temp := res.Next
		res.Next = pre
		pre = res
		res = temp
	}
	return pre
}

func main()  {
	//test := &ListNode{1,&ListNode{2,&ListNode{3,nil,}}}
	//test := &ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,&ListNode{5,nil,}}}}}
	test := &ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,&ListNode{5,&ListNode{6,nil,}}}}}}

	reorderList(test)

	for test != nil {
		fmt.Println(test.Val)
		test = test.Next
	}
}
