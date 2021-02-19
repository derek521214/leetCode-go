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

func reorderList2(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

/**
执行用时为 4 ms 的范例
 */
func reorderList1(head *ListNode)  {
	if head==nil||head.Next==nil{
		return
	}
	fast:=head
	slow:=head
	for fast.Next!=nil&&fast.Next.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	}
	head2:=slow.Next
	slow.Next=nil
	head2=reverse(head2)
	dummy:=&ListNode{}
	cur:=dummy
	cur1:=head
	cur2:=head2
	for cur1!=nil&&cur2!=nil{
		cur.Next=cur1
		cur=cur.Next
		cur1=cur1.Next

		cur.Next=cur2
		cur=cur.Next
		cur2=cur2.Next
	}
	if cur1!=nil{
		cur.Next=cur1
	}
}
func reverse(head*ListNode)*ListNode{
	if head==nil||head.Next==nil{
		return head
	}
	rst:=reverse(head.Next)
	head.Next.Next=head
	head.Next=nil
	return rst
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
