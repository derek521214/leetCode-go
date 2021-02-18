package main

import "fmt"
/**
88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

示例 1：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]

提示：
nums1.length == m + n
nums2.length == n
0 <= m, n <= 200
1 <= m + n <= 200
-109 <= nums1[i], nums2[i] <= 109
 */

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j := 0, 0
	temp := make([]int, 0)
	for  i < m  || j < n  {
		if i == m {
			temp = append(temp, nums2[j:n]... )
			break
		}
		if j == n {
			temp = append(temp, nums1[i:m]... )
			break
		}
		if nums1[i] < nums2[j] {
			temp = append(temp, nums1[i])
			i++
		} else {
			temp = append(temp, nums2[j])
			j++
		}
	}
	fmt.Println(temp)
	copy(nums1, temp)
}

/**
范例
自己理解 是从最后比较大的放到最后 依次直到一个结束 然后把剩下的一次
 */
func merge1(nums1 []int, m int, nums2 []int, n int)  {
	tail1 := m - 1
	tail2 := n - 1
	tail := m + n -1

	for tail1 >= 0 && tail2 >= 0 {
		if nums1[tail1] < nums2[tail2] {
			nums1[tail] = nums2[tail2]
			tail2 -= 1
		} else {
			nums1[tail] = nums1[tail1]
			tail1 -= 1
		}

		tail -= 1
	}

	/*for tail >= 0 && tail1 >= 0 {
		nums1[tail] = nums1[tail1]
		tail -= 1
		tail1 -= 1
	}*/

	for tail >= 0 && tail2 >= 0 {
		nums1[tail] = nums2[tail2]
		tail -= 1
		tail2 -= 1
	}
}

func main()  {
	num1 := []int{1,2,3,0,0,0}
	num2 := []int{2,5,6}
	merge1(num1, 3, num2, 3)
	fmt.Print(num1)
}