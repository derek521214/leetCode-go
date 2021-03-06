package main

/**
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
 */


func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			temp := node.Next
			node.Next = node.Next.Next
			temp.Next = nil
		} else {
			node = node.Next
		}

	}
	return head
}

func main()  {

}

