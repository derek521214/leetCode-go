package main

import "fmt"

/**
148. 排序链表
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

进阶：
你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

示例 1：

输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：

输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]

提示：
链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}


func merge2(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge2(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}
	return dummyHead.Next
}

/**
自己写的超时了
 */
func sortList1(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	res := &ListNode{Next: head}
	if head.Val > head.Next.Val {
		temp := head.Next.Next
		res.Next = head.Next
		head.Next = temp
		res.Next.Next = head
	}
	curr := res.Next.Next
	pre := res.Next
	depth := 0
	for curr != nil && curr.Next != nil {
		if curr.Val > curr.Next.Val {
			temp := curr.Next.Next
			pre.Next = curr.Next
			curr.Next = temp
			pre.Next.Next = curr
		} else {
			curr = curr.Next
		}
		pre = pre.Next
		depth++
	}
	for depth > 0 {
		res = exchange(res)
		depth--
	}
	return res.Next
}

func exchange(node *ListNode) *ListNode {
	curr := node.Next
	pre := node
	for curr != nil && curr.Next != nil {
		if curr.Val > curr.Next.Val {
			temp := curr.Next.Next
			pre.Next = curr.Next
			curr.Next = temp
			pre.Next.Next = curr
		} else {
			curr = curr.Next
		}
		pre = pre.Next
	}
	return node
}

func main()  {
	test := &ListNode{5,&ListNode{3,&ListNode{4,&ListNode{1,nil}}}}
	res := sortList(test)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
