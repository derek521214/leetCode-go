package main

/**
https://leetcode-cn.com/problems/validate-binary-search-tree/
98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
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
func isValidBST(root *TreeNode) bool {
	isValid := true
	min := getMin(root)
	max := getMax(root)
	compare(root.Left, min, root.Val, &isValid)
	if !isValid {
		return false
	}
	compare(root.Right, root.Val, max, &isValid)
	return isValid
}

func compare(node *TreeNode, min, max int, isValid *bool)  {
	if node == nil || !*isValid {
		return
	}
	if node.Val <= min || node.Val >= max {
		*isValid = false
		return
	}
	if node.Left != nil {
		compare(node.Left, min, node.Val, isValid)
	}
	if node.Right != nil {
		compare(node.Right, node.Val, max, isValid)
	}
}

func getMin(node *TreeNode) int  {
	l := node
	for l.Left != nil {
		l = l.Left
	}
	return l.Val - 1
}

func getMax(node *TreeNode) int  {
	r := node
	for r.Right != nil {
		r = r.Right
	}
	return r.Val + 1
}

/**
范例
 */
func isValidBST1(root *TreeNode) bool {
	res := []int{}
	midTran(root, &res)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

func midTran(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	midTran(root.Left, res)
	*res = append(*res, root.Val)
	midTran(root.Right, res)
}