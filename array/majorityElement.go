package main

/**
169. 多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：

输入：[3,2,3]
输出：3
示例 2：

输入：[2,2,1,1,1,2,2]
输出：2

进阶：

尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
 */
func majorityElement(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	mapLen := make(map[int]int)
	for i:=0; i<length; i++ {
		if _, ok := mapLen[nums[i]]; !ok {
			mapLen[nums[i]] = 1
		} else {
			mapLen[nums[i]] += 1
		}

		if mapLen[nums[i]] > length/2 {
			return nums[i]
		}
	}
	return 0
}


func majorityElement1(nums []int) int {
	// 核心思想：两两抵消，最后剩余的数就是结果
	num:=nums[0]
	count:=1
	for i:=1;i<len(nums);i++{
		if nums[i] == num{
			count++
		}else {
			if count-1 == 0{
				i++
				num = nums[i]
				count = 1
			}else{
				count--
			}
		}
	}
	return num
}

