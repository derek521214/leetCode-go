package main

/**
https://leetcode-cn.com/problems/diameter-of-binary-tree/
543. 二叉树的直径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	depTree(root, &res)
	return res
}

func depTree(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	l := depTree(node.Left, res)
	r := depTree(node.Right, res)
	if *res < l + r {
		*res = l + r
	}
	return  max(l, r) + 1
}

func max(l, r int) int  {
	if l > r {
		return l
	}
	return r
}