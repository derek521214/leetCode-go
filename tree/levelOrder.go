package main

import (
	"fmt"
)
/**
102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
 */
type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }

 /**
 深度优先
  */
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	result = depth(root, 0, result)
	return result
}

func depth( node *TreeNode, level int, nums [][]int) [][]int  {
	if node == nil {
		return nums
	}
	if len(nums) == level {
		nums = append(nums, []int{})
	}
	nums[level] = append(nums[level], node.Val)
	nums = depth( node.Left, level+1, nums)
	nums = depth( node.Right, level+1, nums)
	return nums
}



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
广度优先
 */
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result, levelList := [][]int{[]int{root.Val}}, []*TreeNode{root}
	for len(levelList) > 0 {
		var temp []int
		var tempList []*TreeNode
		for _, node := range levelList {
			if node.Left != nil {
				temp = append(temp, node.Left.Val)
				tempList = append(tempList, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right.Val)
				tempList = append(tempList, node.Right)
			}
		}
		if len(temp) > 0 {
			result = append(result, temp)
		}
		levelList = tempList
	}
	return result
}

func main()  {
	test := &TreeNode{
		3,
		&TreeNode{
			9,
			nil,
			nil,
		},
		&TreeNode{
			20,
			&TreeNode{15,nil,nil},
			&TreeNode{7,nil,nil},
		},
	}
	res := levelOrder(test)
	fmt.Println(res)
}

