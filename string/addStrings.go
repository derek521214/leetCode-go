package main

import (
	"fmt"
)

/**
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

提示：
num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
 */

func addStrings1(s1, s2 string) string  {
	len1, len2 := len(s1), len(s2)
	var result string
	if len1 >= len2 {
		add := 0
		m := len2 -1
		for i:= len1 -1; i>=0 ; i-- {
			if m >= 0 {
				num1 := int(s1[i] - '0')
				num2 := int(s2[m] - '0')
				num := num1 + num2 + add
				add = num / 10
				result = fmt.Sprintf("%v", num % 10) + result

				m--
			} else {
				num1 := int(s1[i] - '0')
				num := num1  + add
				add = num / 10
				result =  fmt.Sprintf("%v", num % 10) +result
			}
		}
		if add >0 {
			result =  fmt.Sprintf("%v", add) + result
		}
	} else {

	}
	return result
}

func addStrings(s1, s2 string) string  {
	add := 0
	var result string
	for i, m := len(s1)-1, len(s2)-1; i >=0 || m >= 0 || add !=0; i,m = i-1,m-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(s1[i] - '0')
		}
		if m >= 0 {
			y = int(s2[m] - '0')
		}
		num := x + y + add
		add = num / 10
		result = fmt.Sprintf("%v", num % 10) + result
	}
	return result
}

func main()  {
	s2 := "1"
	s1 := "1231"
	s := addStrings(s1, s2)
	fmt.Print(s)
}
