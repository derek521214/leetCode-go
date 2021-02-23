package main
/**
199. 二叉树的右视图
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
通过次数85,748提交次数132,054
 */


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result, levelList := []int{root.Val},[]*TreeNode{root}
	for len(levelList) > 0 {
		tempList := []*TreeNode{}
		for _, node := range levelList {
			if node.Right != nil {
				tempList = append(tempList, node.Right)
			}
			if node.Left != nil {
				tempList = append(tempList, node.Left)
			}
		}
		if len(tempList) > 0 {
			result = append(result, tempList[0].Val)
		}
		levelList = tempList
	}
	return result
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView1(root *TreeNode) []int {
	ans := []int{}
	layer := 1
	return dfs(root, layer, ans)
}

func dfs(root *TreeNode, layer int, res []int) []int {
	if root == nil {
		return res
	}

	if len(res) < layer {
		res = append(res, root.Val)
	}

	res = dfs(root.Right, layer+1, res)
	res = dfs(root.Left, layer+1, res)
	return res
}