package main

import "fmt"

/**
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"
 */
func longestPalindrome(s string) string {
	slen := len(s)
	if slen < 2 {
		return s
	}
	maxLen := 1
	res := s[0:1]
	for i:= 0; i < slen-1; i++ {
		leftStr := palindrome(s, i, i)
		rightStr := palindrome(s, i, i+1)
		temp :=""
		if len(leftStr) > len(rightStr) {
			temp = leftStr
		} else {
			temp = rightStr
		}
		if len(temp) > maxLen {
			res = temp
			maxLen = len(temp)
		}
	}
	return res
}

func palindrome(s string, left, right int) string  {
	slen := len(s)
	for left >= 0 && right < slen {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return s[left+1:right]
}


/**
暴力破解
 */
func longestPalindrome1(s string) string {
	slen := len(s)
	if slen < 2 {
		return s
	}
	begin := 0
	maxLen := 1
	for i:= 0; i < slen-1; i++ {
		for j:= i+1; j < slen; j++ {
			if j - i + 1 > maxLen && isPalindrome(s[i:j+1]) {
				begin = i
				maxLen = j - i + 1
 			}
		}
	}
	return s[begin:begin+maxLen]
}

func isPalindrome(s string) bool  {
	slen := len(s)
	for i:= 0; i<=slen-1; i++ {
		if s[i] != s[slen-1] {
			return false
		}
		slen--
	}
	return true
}

func main()  {
	//s := "babad"
	s := "bb"
	res := longestPalindrome(s)
	fmt.Print(res)
}