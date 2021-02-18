package main

import "fmt"

/**
94. 二叉树的中序遍历
给定一个二叉树的根节点 root ，返回它的 中序 遍历。

示例 1：

输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：

输入：root = [1,2]
输出：[2,1]
示例 5：

输入：root = [1,null,2]
输出：[1,2]

提示：
树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
 */

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