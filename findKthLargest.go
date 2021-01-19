package main

import "fmt"

/**
215. 数组中的第K个最大元素
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:
你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 */
func findKthLargest(nums []int, k int) int {
	if k >=1 && k <= len(nums) {
		sortNums := sort(nums)
		return sortNums[k-1]
	}
	return 0
}

func sort1(nums []int) []int {
	numLen := len(nums)
	if numLen == 0 {
		return nil
	}
	if numLen == 1 {
		return nums
	} else {
		curr := nums[0]
		var left []int
		var right []int
		for i:=1;i<numLen;i++ {
			if nums[i] > curr {
				left = append(left, nums[i])
			} else {
				right = append(right, nums[i])
			}
		}
		var result []int
		result = append(result, sort1(left)...)
		result = append(result, curr)
		result = append(result, sort1(right)...)
		return result
	}
}

func sort(nums []int) []int  {
	return quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(nums, left, right)
		quickSort(nums, left, partitionIndex-1)
		quickSort(nums, partitionIndex+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	curr := left
	index := curr + 1
	for i := index; i <= right; i++ {
		if nums[i] > nums[curr] {
			swap(nums, i, index)
			index++
		}
	}
	swap(nums, curr, index-1)
	return index - 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main()  {
	nums := []int{3,2,3,1,2,4,5,5,6}
	result := findKthLargest(nums, 4)
	fmt.Println(result)
}