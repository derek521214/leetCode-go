package main

/**
958. 二叉树的完全性检验
给定一个二叉树，确定它是否是一个完全二叉树。

百度百科中对完全二叉树的定义如下：
若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。（注：第 h 层可能包含 1~ 2h 个节点。）

示例 1：

输入：[1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。
示例 2：

输入：[1,2,3,4,5,null,7]
输出：false
解释：值为 7 的结点没有尽可能靠向左侧。

提示：

树中将会有 1 到 100 个结点。
 */

/* type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	end := false
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr_node := queue[0]
		queue = queue[1:]
		if curr_node == nil {
			end = true
		} else {
			if end == true {
				return false
			}
			queue = append(queue, curr_node.Left)
			queue = append(queue, curr_node.Right)
		}
	}
	return true
}

var ids map[int]bool

func isCompleteTree1(root *TreeNode) bool {
	ids = make(map[int]bool)
	// 根节点的下标是0
	dfs(root, 0)
	for i := 0; i < len(ids); i++ {
		if !ids[i] {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, id int) {
	if node == nil {
		return
	}
	ids[id] = true
	// 父节点的下标是id，那么左右儿子的下标分别是id*2+1和id*2+2
	id <<= 1
	dfs(node.Left, id + 1)
	dfs(node.Right, id + 2)
}