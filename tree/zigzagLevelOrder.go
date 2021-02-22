package main

import "fmt"

/**
103. 二叉树的锯齿形层序遍历
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
通过次数115,194提交次数201,838
 */

 type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	res = append(res,[]int{root.Val})
	ceng := []*TreeNode{root}
	depth := 0
	for len(ceng) > 0 {
		temp := []int{}
		tempList := []*TreeNode{}
		if depth % 2 == 1 {
			for i:= len(ceng)-1; i>=0;i-- {
				if ceng[i].Left != nil {
					temp = append(temp, ceng[i].Left.Val)
					tempList = append(tempList, ceng[i].Left)
				}
				if ceng[i].Right != nil {
					temp = append(temp, ceng[i].Right.Val)
					tempList = append(tempList, ceng[i].Right)
				}
			}
		} else {
			for i:= len(ceng)-1; i>=0;i-- {
				if ceng[i].Right != nil {
					temp = append(temp, ceng[i].Right.Val)
					tempList = append(tempList, ceng[i].Right)
				}
				if ceng[i].Left != nil {
					temp = append(temp, ceng[i].Left.Val)
					tempList = append(tempList, ceng[i].Left)
				}
			}
		}
		if len(temp) > 0 {
			res = append(res, temp)
		}
		ceng = tempList
		depth++
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result = nil
	l2rScanTree(root, 0)
	r2lScanTree(root, 0)
	return result
}

var result [][]int

func l2rScanTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if len(result) == level {
		result = append(result, []int{})
	}
	if level%2 == 0 {
		result[level] = append(result[level], root.Val)
	}
	l2rScanTree(root.Left, level+1)
	l2rScanTree(root.Right, level+1)
}

func r2lScanTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level%2 == 1 {
		result[level] = append(result[level], root.Val)
	}
	r2lScanTree(root.Right, level+1)
	r2lScanTree(root.Left, level+1)
}


func main()  {
	test := &TreeNode{
		1,
			&TreeNode{
				2,
					&TreeNode{4,nil,nil},
					nil,
			},
			&TreeNode{
				3,
					nil,
					&TreeNode{5,nil,nil},
			},
	}
	res := zigzagLevelOrder(test)
	fmt.Println(res)
}