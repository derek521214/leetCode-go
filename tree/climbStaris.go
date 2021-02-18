package main

import "fmt"

/**
70. 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
 */

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	 nums := make([]int, n+1)
	for i:=0; i<= n; i++ {
		if i==0 {
			nums[i] = 0
		}
		if i == 1 {
			nums[i] = 1
		}
		if i == 2 {
			nums[2] = 2
		}
		if i>2 {
			nums[i] = nums[i-1] + nums[i-2]
		}
	}
	return nums[n]
}

func climbStairs1(n int) int {
	one, two, result := 0, 0, 1
	for i := 1; i <= n; i++ {
		one = two
		two = result
		result = one + two
	}
	return result
}

func main()  {
	res := climbStairs(3)
	fmt.Print(res)
}
