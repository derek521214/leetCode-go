package main

import "fmt"

/**
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]

提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	var mapIndex  []*ListNode
	curr := head
	for curr != nil {
		mapIndex = append(mapIndex, curr)
		curr = curr.Next
	}
	listLen := len(mapIndex)
	if n == 1 {
		mapIndex[listLen-2].Next = nil
	} else if listLen == n {
		head = mapIndex[1]
		mapIndex[0] = nil
	} else {
		mapIndex[listLen-n-1].Next = mapIndex[listLen-n+1]
		mapIndex[listLen-n].Next = nil
	}
	return head
}

/**
执行用时为 0 ms 的范例
 */
func removeNthFromEnd1(head *ListNode, n int) *ListNode  {
	dumpty := &ListNode{Next: head}
	fast := head
	for ; n>0 ; n-- {
		fast = fast.Next
	}
	slow := dumpty
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return  dumpty.Next
}

/**
执行消耗内存为 2288 kb 的范例
 */
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummyNode.Next = head

	length := 0
	p := head
	for p != nil {
		length++
		p = p.Next
	}

	deleteIndex := length - n
	p = dummyNode
	index := 0
	for index != deleteIndex {
		p = p.Next
		index++
	}

	// for i := 0; i < length-n; i++ {
	// 	p = p.Next
	// }
	p.Next = p.Next.Next

	return dummyNode.Next
}


func main()  {

	test := &ListNode{1,&ListNode{2,&ListNode{3,&ListNode{4,nil,}}}}

	res := removeNthFromEnd1(test, 4)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}