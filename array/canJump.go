package main

import "fmt"

/**
55. 跳跃游戏
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。



示例 1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


提示：

1 <= nums.length <= 3 * 104
0 <= nums[i] <= 105
 */

/**
转换为找0
 */
func canJump(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}
	var zero []int
	//遍历找值为0的索引
	for i:= 0 ; i < length-1; i++ {
		if nums[i] == 0 {
			zero = append(zero, i)
		}
	}
	zlen := len(zero)
	//如果没有
	if zlen == 0 {
		return true
	}
	temp := 0
	for _, index := range zero {
		for i:= 0 ; i < index; i++ {
			if i + nums[i] > index {
				temp++
				break
			}
		}
	}
	if temp == zlen {
		return true
	}
	return false
}

/**
执行用时为 0 ms 的范例
 */
func canJump1(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums); i++ {
		if i > maxPos {
			return false
		}
		if nums[i]+i > maxPos {
			maxPos = nums[i] + i
			if maxPos >= len(nums)-1 {
				return true
			}
		}

	}
	return true
}

/**
执行消耗内存为 4048 kb 的范例
 */
func canJump2(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}
	if len(nums) > 1 && nums[0] == 0 {
		return false
	}
	if nums[0]==len(nums)-1{
		return true
	}
	jump := false
	for i := 0; i < len(nums); {
		if nums[i] != 0 {
			i++
			continue
		}
		if i == len(nums)-1 {
			return true
		}
		for j := 0; j < i; j++ {
			if nums[j]+j > i {
				jump = true
				break
			}
		}
		if jump {
			i++
			jump = false
			continue
		}else{
			return false
		}
	}
	return true
}

func main()  {
	test :=[]int{3,0,8,2,0,0,1}

	fmt.Println(canJump(test))

}