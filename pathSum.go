package main

import "fmt"

/**
https://leetcode-cn.com/problems/path-sum-ii/
113. 路径总和 II
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
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
// 用于返回
var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	// 如果root为空那就直接滚蛋
	if root==nil{
		return nil
	}
	// 把root对应的值先加入到当前计数
	now:=[]int{root.Val}
	nowSum:=root.Val
	// 初始化一下返回值
	res=make([][]int,0)
	// dfs
	dfs(root,now,nowSum,sum)
	return res
}
/**
node:当前dfs到的节点
now :根节点到当前的路径
sum ：根节点到当前的路径和
target：目标和
*/

func dfs (node *TreeNode,now []int,sum,target int){
	// 结束条件
	// 左右节点都不为空（叶子节点），那必须得结束了
	if node.Left==nil && node.Right==nil{
		// 结束的时候 sum当前路径和==目标和
		// 把当前路径追加到结果
		if sum==target{
			temp:=make([]int,len(now))
			copy(temp,now)
			res=append(res,temp)
		}
		// 不管有没有追加都要return回上层了
		return
	}
	// 如果左节点不为空，左递归
	if node.Left!=nil{
		// 把左节点加入now 和 sum
		now=append(now,node.Left.Val)
		sum=sum+node.Left.Val
		dfs(node.Left,now,sum,target)
		// 回溯：回来之后把节点在now 和 sum 里摘除
		now=now[:len(now)-1]
		sum=sum-node.Left.Val
	}
	// 右节点递归 同理
	if node.Right!=nil{
		now=append(now,node.Right.Val)
		sum=sum+node.Right.Val
		dfs(node.Right,now,sum,target)
		now=now[:len(now)-1]
		sum=sum-node.Right.Val
	}
	return
}

func main()  {
	test := &TreeNode{
		Val: -2,
		Right: &TreeNode{
			Val: -3,
		},
	}
	res := pathSum(test, -5)
	fmt.Print(res)
}