package main

import "fmt"

/**
206. 反转链表
反转一个单链表。
示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val	int
	Next *ListNode
}

/**
解题思路
首先定义一个cur指针，指向头结点，再定义一个pre指针，初始化为null。
然后就要开始反转了，首先要把 cur->next 节点用tmp指针保存一下，也就是保存一下这个节点。
为什么要保存一下这个节点呢，因为接下来要改变 cur->next 的指向了，将cur->next 指向pre ，此时已经反转了第一个节点了。
接下来，就是循环走如下代码逻辑了，继续移动pre和cur指针。
最后，cur 指针已经指向了null，循环结束，链表也反转完毕了。 此时我们return pre指针就可以了，pre指针就指向了新的头结点。
作者：carlsun-2
链接：https://leetcode-cn.com/problems/reverse-linked-list/solution/206-fan-zhuan-lian-biao-shuang-zhi-zhen-fa-di-gui-/
来源：力扣（LeetCode）
 */
func reverseList1(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre *ListNode, curr *ListNode) *ListNode {
	if curr == nil {
		return pre
	}
	temp := curr.Next
	curr.Next = pre
	return reverse(curr, temp)
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
	newList := reverseList(list1)
	fmt.Println(newList)
	for newList != nil {
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}