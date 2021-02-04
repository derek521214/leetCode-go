package main
/**
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

进阶：
你可以运用递归和迭代两种方法解决这个问题吗？
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	isSym := true
	check(root, root, &isSym)
	return isSym
}

func check(p, q *TreeNode, isSym *bool)  {
	if p == nil && q == nil {
		return
	}
	if p == nil || q == nil {
		*isSym = false
		return
	}
	if p.Val != q.Val {
		*isSym = false
		return
	}
	check(p.Left, q.Right, isSym)
	check(p.Right, q.Left, isSym)
}

func isSymmetric1(root *TreeNode) bool {
	return check1(root, root)
}

func check1(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check1(p.Left, q.Right) && check1(q.Right, p.Left)
}

