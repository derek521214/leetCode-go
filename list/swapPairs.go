package main

import "fmt"

/**
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：

输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]

提示：

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := &ListNode{
		Next: head,
	}
	curr := temp
	for curr.Next != nil && curr.Next.Next != nil {
		////最后的头指针
		//last := curr.Next.Next.Next
		//下一个
		next := curr.Next
		curr.Next = curr.Next.Next
		next.Next = curr.Next.Next
		curr.Next.Next = next

		curr = curr.Next.Next
	}
	return temp.Next
}

func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	tmp := dummy
	for tmp.Next != nil && tmp.Next.Next != nil {
		slow := tmp.Next
		fast := tmp.Next.Next
		tmp.Next = fast
		slow.Next = fast.Next
		fast.Next = slow
		tmp = slow
	}
	return dummy.Next
}

/**
执行消耗内存为 2100 kb 的范例
 */
func swapPairs2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead
	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil {
		p1.Next = p2.Next
		p2.Next = p1
		pre.Next = p2

		//p2,p1
		pre = p1
		p1 = p1.Next
		if p1!= nil {
			p2 = p1.Next
		}
	}

	return dummyHead.Next
}

func main()  {
	test := &ListNode{
		1,&ListNode{2,&ListNode{3,&ListNode{4,nil}}},
	}
	res := swapPairs(test)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}