package main

import "fmt"

/**
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

示例 1:

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。
示例 2:

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
示例 3:

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。

提示:

1 <= text1.length <= 1000
1 <= text2.length <= 1000
输入的字符串只含有小写英文字符。
 */

func longestCommonSubsequence(text1 string, text2 string) int {
	t1Len, t2Len := len(text1), len(text2)
	if t1Len == 0 || t2Len == 0 {
		return 0
	}
	arr := make([][]int, t1Len + 1)
	arr[0] = make([]int, t2Len+1)
	for i:=1; i <= t1Len; i++ {
		arr[i] = make([]int, t2Len+1)
		for j :=1; j<=t2Len; j++ {
			if text1[i-1] == text2[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else {
				arr[i][j] = max(arr[i-1][j], arr[i][j-1])
			}
		}
	}
	return arr[t1Len][t2Len]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main()  {
	t1 := "ezupkr"
	t2 := "ubmrapg"
	res := longestCommonSubsequence(t1, t2)
	fmt.Println(res)
}