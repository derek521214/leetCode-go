package main

import "fmt"

/**
1358. 包含所有三种字符的子字符串数目
给你一个字符串 s ，它只包含三种字符 a, b 和 c 。

请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。

示例 1：

输入：s = "abcabc"
输出：10
解释：包含 a，b 和 c 各至少一次的子字符串为 "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" 和 "abc" (相同字符串算多次)。
示例 2：

输入：s = "aaacb"
输出：3
解释：包含 a，b 和 c 各至少一次的子字符串为 "aaacb", "aacb" 和 "acb" 。
示例 3：

输入：s = "abc"
输出：1


提示：

3 <= s.length <= 5 x 10^4
s 只包含字符 a，b 和 c 。
 */

func numberOfSubstrings(s string) int {
	total,left,numA,numB,numC := 0,0,0,0,0
	strlen := len(s)
	for i:=0; i<strlen; i++ {
		if s[i] == 'a' {
			numA++
		} else if s[i] == 'b' {
			numB++
		} else {
			numC++
		}
		curr :=0
		for numA > 0 && numB > 0 && numC > 0 {
			curr++
			if s[left] == 'a' {
				numA--
			} else if s[left] == 'b' {
				numB--
			} else {
				numC--
			}
			left++
		}
		total += curr*(strlen-i)
	}
	return total
}

/**
提交后超时 代码逻辑没问题
 */
func numberOfSubstrings1(s string) int {
	total := 0
	depthAdd(s, &total)
	return total
}

func depthAdd(s string, total *int)  {
	strLen := len(s)
	if strLen < 3 {
		return
	}
	mapStr := make(map[uint8]struct{}, 3)
	num := 0
	for i:=0; i< strLen; i++ {
		if _, ok := mapStr[s[i]]; !ok {
			mapStr[s[i]] = struct{}{}
		}
		if len(mapStr) == 3 {
			num = strLen-i
			break
		}
	}
	if num > 0 {
		*total = *total + num
		depthAdd(s[1:], total)
	} else {
		return
	}
}

func main()  {
	test := "abcabc"
	//test := "aaacb"
	res := numberOfSubstrings(test)
	fmt.Println(res)
}