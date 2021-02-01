package main

import "fmt"

/**
151. 翻转字符串里的单词
给定一个字符串，逐个翻转字符串中的每个单词。

说明：

无空格字符构成一个 单词 。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 */

func reverseWords(s string) string {
	slen := len(s)
	if slen < 2 {
		return s
	}
	strSlice := []string{}
	start := 0
	for i:=0; i< slen; i++ {
		if s[i] == ' ' {
			if i > start {
				strSlice = append(strSlice, s[start:i])
			}
			start = i+1
		}
		if i == slen-1 && i >= start {
			strSlice = append(strSlice, s[start:slen])
		}
	}
	sslen := len(strSlice)
	if sslen >= 1 {
		res := ""
		for j:= sslen-1; j>=0 ; j-- {
			if j > 0 {
				res += strSlice[j] + " "
			} else {
				res += strSlice[j]
			}
		}
		return res
	}
	return ""
}

/**
范例
 */

func reverseWords1(s string) string {
	res := make([]byte, len(s))

	lc := byte(' ') // last visited char
	j, k := 0, len(s) - 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if lc == ' ' { // consective empty char
				continue
			}

			for k < len(s)-1 {
				res[j] = res[k+1]
				j++
				k++
			}

			res[j] = ' '
			j++
			k = len(s) - 1
		}
		lc = s[i]

		if s[i] != ' ' {
			res[k] = s[i]
			k--
		}
	}

	for k < len(s)-1 {
		res[j] = res[k+1]
		j++
		k++
	}

	if j < len(s) {
		res = res[:j]
	}

	if res[j - 1] == ' ' {
		res = res[:j-1]
	}

	return string(res)
}

func main()  {
	str := "the sky is blue"
	//str := " asdasd df f"
	res :=reverseWords(str)
	fmt.Print(res)
}
