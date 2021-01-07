package main

import "fmt"
/**
寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000

提示：
nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenNum1, lenNum2 := len(nums1), len(nums2)
	var result float64
	if lenNum1 + lenNum2 >= 1 {
		var merge []int
		index1 := 0
		index2 := 0
		for index1 < lenNum1 || index2 < lenNum2 {
			if index1 == lenNum1 {
				merge = append(merge, nums2[index2])
				index2++
			} else if index2 == lenNum2 {
				merge = append(merge, nums1[index1])
				index1++
			} else if nums1[index1] > nums2[index2] {
				merge = append(merge, nums2[index2])
				if index2 < lenNum2 {
					index2++
				}
			} else {
				merge = append(merge, nums1[index1])
				if index1 < lenNum1 {
					index1++
				}
			}
		}
		total := len(merge)
		if total == 1 {
			return float64(merge[0])
		}
		if total % 2 == 1 {
			return float64(merge[total/2])
		} else {
			return (float64(merge[total/2]) + float64(merge[(total/2)-1]))/float64(2)
		}
	}
	return result
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	lenNum1, lenNum2 := len(nums1), len(nums2)
	var result float64
	if lenNum1 + lenNum2 >= 1 {
		merge := make([]int, lenNum1 + lenNum2)
		index1 := 0
		index2 := 0
		index := 0
		for index1 < lenNum1 || index2 < lenNum2 {
			if index1 == lenNum1 {
				merge[index] = nums2[index2]
				index2++
			} else if index2 == lenNum2 {
				merge[index] = nums1[index1]
				index1++
			} else if nums1[index1] > nums2[index2] {
				merge[index] = nums2[index2]
				if index2 < lenNum2 {
					index2++
				}
			} else {
				merge[index] = nums1[index1]
				if index1 < lenNum1 {
					index1++
				}
			}
			index++
		}
		total := len(merge)
		if total == 1 {
			return float64(merge[0])
		}
		if total % 2 == 1 {
			return float64(merge[total/2])
		} else {
			return (float64(merge[total/2]) + float64(merge[(total/2)-1]))/float64(2)
		}
	}
	return result
}

func main()  {
	num1 := []int{1,2,3}
	num2 := []int{2,3,4}
	res :=findMedianSortedArrays1(num1, num2)
	fmt.Println(res)
}
