package main

import "fmt"

/**
两数相加
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
 

提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{
		Val:  0,
		Next: nil,
	}
	result := getListNode(l1, l2, l3,0)
	return result
}

func getListNode(l1 *ListNode, l2 *ListNode, l3 *ListNode, num int) *ListNode {
	if (l1 != nil || l2 != nil) && (l3 != nil && l3.Next == nil) {
		total := 0
		next := 0
		if l1 == nil {
			total = l2.Val + num
		} else if l2 == nil {
			total = l1.Val + num
		} else {
			total = l1.Val + l2.Val + num
		}
		if total > 10 {
			l3.Val = total - 10
			next = 1
		} else {
			l3.Val = total
		}
		if l1 == nil {
			l3.Next = getListNode(nil, l2.Next, &ListNode{
				Val:  0,
				Next: nil,
			}, next)
		} else if l2 == nil {
			l3.Next = getListNode(l1.Next, nil, &ListNode{
				Val:  0,
				Next: nil,
			}, next)
		} else {
			l3.Next = getListNode(l1.Next, l2.Next, &ListNode{
				Val:  0,
				Next: nil,
			}, next)
		}
	} else {
		if num == 0 {
			return nil
		}
		return &ListNode{
			Val:  num,
			Next: nil,
		}
	}
	return l3
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	next := 0
	cur := result
	for l1 != nil || l2 != nil {
		num1, num2 := 0, 0
		if l1 != nil {
			num1 = l1.Val
		}
		if l2 != nil {
			num2 = l2.Val
		}
		total := num1 + num2 + next
		if total >= 10 {
			next = 1
			cur.Next = &ListNode{
				Val:  total - 10,
				Next: nil,
			}
		} else {
			next = 0
			cur.Next = &ListNode{
				Val:  total,
				Next: nil,
			}
		}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if next > 0 {
		cur.Next = &ListNode{
			Val:  next,
			Next: nil,
		}
	}
	return result.Next
}

func main()  {
	l1 := &ListNode{
		Val:  9,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{
				Val:  9,
				Next: &ListNode{
					Val:  9,
					Next: &ListNode{
						Val:  9,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
		},
	}

	l2 := &ListNode{
		Val:  9,
		Next: &ListNode{
			Val:  9,
			Next: &ListNode{
				Val:  9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	res := addTwoNumbers(l1, l2)

	if res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	fmt.Println(res)
}