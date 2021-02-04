package main

import "fmt"
import "sort"

/**
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：
1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
 */



func merge(intervals [][]int) [][]int {
	intLen := len(intervals)
	if intLen <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
 	start := intervals[0][0]
 	end := intervals[0][1]
 	fmt.Println(start, end)
	var result [][]int
 	for i:=1; i <len(intervals); i++ {
 		if intervals[i][0] > end {
 			result = append(result, []int{start, end})
 			start = intervals[i][0]
 			end = intervals[i][1]
		} else if intervals[i][1] > end {
			end = intervals[i][1]
		}
		if i == len(intervals)-1 {
			result = append(result, []int{start, end})
		}
	}
	return result
}

/**

 */
func merge1(intervals [][]int) [][]int {
	// 有合并，则合并
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	output := make([][]int,1,len(intervals))
	output[0] = intervals[0]
	for _, v := range intervals[1:] {
		flag := true
		for ii, iv := range output {
			if v[0] <= iv[1] {
				flag = false
				if v[1] > iv[1] {
					output[ii][1] = v[1]
				}
			}
		}
		if flag {
			output = append(output,v)
		}

	}
	// 没合并，则返回
	return output
}

/**
执行消耗内存为 4612 kb 的范例
 */
func merge2(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	cnt := 1
	for i := 1; i < len(intervals); i++ {
		//逐个折半查找，插入排序，插入时进行合并
		min := 0
		max := cnt - 1
		for min <= max {
			med := (min + max) / 2
			if intervals[i][0] > intervals[med][1] {
				//当前区间在med区间右侧
				min = med + 1
				continue
			}
			if intervals[i][1] < intervals[med][0] {
				//当前区间在med区间左侧
				max = med - 1
				continue
			}
			//当前区间与med区间重叠
			j := med + 1
			k := med + 1
			merged := 0
			if intervals[i][0] < intervals[med][0] {
				//当前区间左侧小于med区间，向左扩展med区间并尝试合并更左侧的区间
				intervals[med][0] = intervals[i][0]
				for j = med - 1; j >= 0; j-- {
					if intervals[med][0] > intervals[j][1] {
						//与左侧区间不重叠
						break
					}
					//与左侧区间重叠，合并区间
					if intervals[j][0] < intervals[med][0] {
						intervals[med][0] = intervals[j][0]
					}
				}
			}
			if intervals[i][1] > intervals[med][1] {
				//当前区间右侧大于med区间，向右扩展med区间并尝试合并更右侧的区间
				intervals[med][1] = intervals[i][1]
				for k = med + 1; k < cnt; k++ {
					if intervals[med][1] < intervals[k][0] {
						//与右侧区间不重叠
						break
					}
					//与右侧区间重叠，合并区间
					if intervals[k][1] > intervals[med][1] {
						intervals[med][1] = intervals[k][1]
					}
				}
			}
			if j < med - 1 {
				//合并了左侧区间，将med左移到被合并的第一个区间位置
				intervals[j + 1][0] = intervals[med][0]
				intervals[j + 1][1] = intervals[med][1]
				merged += med - 1 - j
				j += 2
			} else {
				j = med + 1
			}
			if k > med + 1 || merged > 0 {
				//合并了右侧区间，将右侧剩余区间移动到med后
				merged += k - med - 1
				for k < cnt {
					intervals[j][0] = intervals[k][0]
					intervals[j][1] = intervals[k][1]
					j++
					k++
				}
			}
			cnt -= merged
			break
		}
		if min > max {
			//没有重叠的有序区间，将i区间插入min位置
			if i != min {
				v0 := intervals[i][0]
				v1 := intervals[i][1]
				for j := cnt - 1; j >= min; j-- {
					intervals[j+1][0] = intervals[j][0]
					intervals[j+1][1] = intervals[j][1]
				}
				intervals[max+1][0] = v0
				intervals[max+1][1] = v1
			}
			cnt++
		}
	}

	return intervals[:cnt]
}

func main()  {
	test := [][]int {
		{1,4},
		{2,3},
	}
	res := merge(test)
	fmt.Print(res)
}
