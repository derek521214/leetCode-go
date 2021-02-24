package main

import "fmt"

/**
105. 从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
通过次数155,914提交次数225,407
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	result := &TreeNode{}
	depth(result, preorder, inorder)
	return result
}

func depth(node *TreeNode, preorder []int, inorder []int)  {
	if len(preorder) == 0 {
		return
	}
	node.Val = preorder[0]
	//根节点索引
	rootIndex := 0
	for i:=0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			rootIndex = i
			break
		}
	}

	leftPreorder := preorder[1:rootIndex+1]
	leftInorder := inorder[0:rootIndex]
	if len(leftPreorder) > 0 {
		node.Left = &TreeNode{}
		depth(node.Left,leftPreorder, leftInorder)
	}

	rightPreorder := preorder[rootIndex+1:]
	rightInorder := inorder[rootIndex+1:]
	if len(rightPreorder) > 0 {
		node.Right = &TreeNode{}
		depth(node.Right,rightPreorder, rightInorder)
	}
 }

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func main()  {

	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}

	res := buildTree(preorder, inorder)

	fmt.Println(res)

}
