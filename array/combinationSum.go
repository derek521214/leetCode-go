package main

import (
	"fmt"
	"sort"
)

/**
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
 */

func combinationSum11(candidates []int, target int) [][]int {
	numLen := len(candidates)
	if numLen == 0 {
		return nil
	}
	if numLen == 1 {
		if candidates[0] == target {
			return [][]int{candidates}
		}
		return nil
	}
	//排序数组
	sort.Ints(candidates)
	//最小值大于目标值
	if candidates[0] > target {
		return nil
	}
	//去掉大于目标值的数据
	if candidates[numLen-1] > target {
		for i:= numLen-2; i>=0; i-- {
			if candidates[i] <= target {
				candidates = candidates[:i+1]
				break
			}
		}
	}

	fmt.Println(candidates)
	return nil
}

/**
作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/combination-sum/solution/zu-he-zong-he-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}

func main()  {
	test := []int{1,3,2,4,5,10,12,13}
	target := 10
	res := combinationSum(test, target)
	fmt.Println(res)
}