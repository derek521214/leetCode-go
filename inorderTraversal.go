package main

import "fmt"

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
func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	iterate(root, &result)
	return  result
}

func iterate(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	iterate(node.Left, nums)
	*nums = append(*nums, node.Val)
	iterate(node.Right, nums)
}

func main()  {
	tree1 := &TreeNode{
		Val: 1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	result := inorderTraversal(tree1)
	fmt.Print(result)
}