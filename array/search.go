package main

import "fmt"

/**
33. 搜索旋转排序数组
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2] ）。
请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4
示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1
示例 3：
输入：nums = [1], target = 0
输出：-1
提示：
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
nums 肯定会在某个点上旋转
-10^4 <= target <= 10^4
*/

func search(nums []int, target int) int {
	nlen := len(nums)
	if nlen == 0 {
		return -1
	}
	if nlen == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	temp := nums[0]-1
	isfirst := true
	for i := 0; i < nlen; i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target  {
			if isfirst {
				temp = nums[i]
				isfirst = false
			}
			if temp > nums[i] {
				return -1
			}
		}
	}
	return -1
}

/**
执行用时为 0 ms 的范例
*/
func search2(nums []int, target int) int {
	l := len(nums)
	lo, hi := 0, l - 1
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		if nums[mid] == target {
			return mid
		} else {

			//左边有序
			if nums[lo] < nums[mid] {
				if target >= nums[lo] && target < nums[mid] {
					hi = mid - 1
				} else {
					lo = mid + 1
				}
				//右边有序
			} else if nums[lo] > nums[mid] {
				if target > nums[mid] && target <= nums[hi] {
					lo = mid + 1
				} else {
					hi = mid - 1
				}
			} else {
				//刚好落在边界上 // 3 1 -> 1
				if nums[lo] > nums[hi] {
					if target < nums[lo] {
						lo = mid + 1
					} else {
						hi = mid - 1
					}
				} else {
					if target > nums[mid] {
						lo = mid + 1
					} else {
						hi = mid - 1
					}
				}
			}
		}
	}
	return -1
}

/**
执行消耗内存为 2604 kb 的范例
*/
func search1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}


func main()  {
	//test := []int{1,3}
	//test := []int{4,5,6,7,0,1,2}
	//test := []int{4,5,6,7,9,1,2,3}
	test := []int{5,1,2,3,4}
	res := search1(test, 1)
	fmt.Print(res)
}