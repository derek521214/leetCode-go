package main

import "fmt"

/**

 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == headB {
		return headA
	}
	h1, h2 := headA, headB
	len1 , len2 := 0, 0
	for h1 != nil || h2 != nil {
		if h1 != nil {
			len1++
			h1 = h1.Next
		}
		if h2 != nil {
			len2++
			h2 = h2.Next
		}
		if h1 != nil && h1 == h2 {
			return h1
		}
	}
	if h1 == h2 {
		if len1 > len2 {
			return findList(headA, headB, len1 - len2)
		} else {
			return findList(headB, headA, len2 - len1)
		}
	}
	return nil
}

func findList(headA, headB *ListNode, reduce int) *ListNode  {
	for headA != nil || headB != nil {
		if reduce ==0 && headB != nil {
			headB = headB.Next
		}
		if headA != nil {
			headA = headA.Next
			if reduce >0 {
				reduce--
			}
		}
		if headA == headB {
			return headA
		}
	}
	return nil
}


func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	mapList := make(map[*ListNode]struct{})
	for headA != nil || headB != nil {
		if headA != nil {
			if _, ok := mapList[headA]; ok {
				return headA
			}
			mapList[headA] = struct{}{}
			headA = headA.Next
		}
		if headB != nil {
			if _, ok := mapList[headB]; ok {
				return headB
			}
			mapList[headB] = struct{}{}
			headB = headB.Next
		}
	}
	return nil
}

func main()  {
	/*l1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next:  &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next:  &ListNode{
				Val: 1,
				Next: nil,
			},
		},
	}
	l2.Next.Next.Next = l1.Next.Next
	res := getIntersectionNode(l1, l2)
	fmt.Print(res)

	l3 := &ListNode{
		Val: 1,
		Next: nil,
	}
	var l4 *ListNode
	l3.Next = l4
	res1 := getIntersectionNode(l3, l4)
	fmt.Print(res1)*/

	l5 := &ListNode{
		Val: 1,
		Next: nil,
	}
	l6 := l5
	res2 := getIntersectionNode(l5, l6)
	fmt.Print(res2)
}

