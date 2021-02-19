package main

/**
234. 回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var nums []int
	depth(head, &nums)
	nlen := len(nums)
	left := 0
	right := nlen-1
	for left <= right {
		if nums[left] != nums[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func depth(node *ListNode, nums *[]int)  {
	for node != nil {
		*nums = append(*nums, node.Val)
		node = node.Next
	}
}

/**
执行用时为 4 ms 的范例
快慢指针 快指针到终点时慢指针到中间 然后慢指针反转（即尾改成头了） 然后再和head依次对比
 */
func isPalindrome1(head *ListNode) bool {
	slow,fast := head,head
	// 快慢指针找到链表的中间节点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	reverse := reverseList(slow)
	for reverse != nil{
		if head.Val != reverse.Val {
			return false
		}
		head = head.Next
		reverse = reverse.Next
	}
	return true
}

// 反转链表
func reverseList(l *ListNode) *ListNode {
	var pre *ListNode
	for l != nil {
		temp := l.Next
		l.Next = pre
		pre = l
		l = temp
	}
	return pre
}

/**
执行消耗内存为 5596 kb 的范例
获取长度 然后跳到中间位置 反转 然后依次对比
 */
func isPalindrome2(head *ListNode) bool {
	if head==nil || head.Next==nil{
		return true
	}
	length, cur := 0, head
	for cur != nil {
		length++
		cur = cur.Next
	}
	sign := length % 2
	length /= 2
	cur = head
	var second *ListNode
	for length > 1 {
		cur = cur.Next
		length--
	}
	if sign == 1 {
		second = cur.Next.Next
		cur.Next.Next = nil
	} else {
		second = cur.Next
		cur.Next = nil
	}
	var dummy *ListNode
	for second != nil {
		mid := second.Next
		second.Next = dummy
		dummy = second
		second = mid
	}
	for head != nil && dummy != nil {
		if head.Val != dummy.Val {
			return false
		}
		head = head.Next
		dummy = dummy.Next
	}
	return true
}