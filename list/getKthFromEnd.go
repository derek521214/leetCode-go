package main

import "fmt"

/**
剑指 Offer 22. 链表中倒数第k个节点
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 2 个节点是值为 4 的节点。

示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.

返回链表 4->5.
 */

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var length int
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	for head != nil && length >= k {
		if length == k {
			return head
		}
		length--
		head = head.Next
	}
	return nil
}

/**
执行用时为 0 ms 的范例
执行消耗内存为 2280 kb 的范例
 */
func getKthFromEnd1(head *ListNode, k int) *ListNode {
	if head == nil || k < 1 {
		return nil
	}
	var slow, fast = head, head
	for i := 1; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			return nil
		}
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func main()  {
	test := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
			3,
			&ListNode{
				4,&ListNode{
					5,nil,
				},
			},
			},
		},
	}
	/*test := &ListNode{
		1,nil,
	}*/

	res := getKthFromEnd(test, 2)
	fmt.Println(res)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}