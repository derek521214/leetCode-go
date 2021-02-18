package main

import "debug/macho"

/**
110. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：true
示例 2：

输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
示例 3：

输入：root = []
输出：true

提示：
树中的节点数在范围 [0, 5000] 内
-104 <= Node.val <= 104
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	res := true
	depth(root, &res)
	return res
}

func depth(node *TreeNode, isBala *bool) int {
	if node == nil || !*isBala {
		return 0
	}
	l := depth(node.Left, isBala)
	r := depth(node.Right, isBala)
	if res := l - r; res > 1 || res < -1 {
		*isBala = false
	}
	return max(l, r) + 1
}

func max(x, y int) int  {
	if x > y {
		return x
	}
	return y
}

/**
范例
 */
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHight := maxDepth(root.Left)
	rightHight := maxDepth(root.Right)
	tmp := leftHight - rightHight
	if tmp > 1 || tmp < (-1) {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}