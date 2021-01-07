package main

import "fmt"
/**
无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0
 
提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func lengthOfLongestSubstring(s string) int {
	lenNum := len(s)
	if lenNum >= 1 {
		max := 0
		for i := 0; i <lenNum; i++ {
			mapString := make(map[string]bool)
			total := 0
			for m := i; m < lenNum; m++ {
				if _, ok := mapString[s[m:m+1]]; !ok {
					mapString[s[m:m+1]] = true
					total = len(mapString)
				} else {
					if total > max {
						max = total
					}
					break
				}
				if m == lenNum -1 && total > max {
					max = total
				}

			}
		}
		return max
	} else {
		return 0
	}
}

func lengthOfLongestSubstring1(s string) int {
	lenNum := len(s)
	if lenNum > 1 {
		sSlice := []byte(s)
		max := 0
		total := 0
		var strSlice []byte
		mapString := make(map[byte]int)
		for i := 0; i < lenNum; i++ {
			if _, ok := mapString[sSlice[i]]; !ok {
				strSlice = append(strSlice, sSlice[i])
				mapString[sSlice[i]] = len(strSlice)-1
				total = len(mapString)
			} else {
				index := mapString[sSlice[i]]
				for m:= 0; m <= index; m++ {
					delete(mapString, strSlice[m])
				}
				i--
				strSlice = strSlice[index+1:]
				for k, v := range strSlice {
					mapString[v] = k
				}
				if total > max {
					max = total
				}
				total = 0
			}
			if i == lenNum - 1 && total > max {
				max = total
			}
		}
		return max
	} else if lenNum==1 {
		return 1
	} else {
		return 0
	}
}

func lengthOfLongestSubstring2(s string) int {
	lenNum := len(s)
	if lenNum > 1 {
		sSlice := []byte(s)
		max := 0
		total := 0
		start := 0
		end := 0
		mapString := make(map[byte]int)
		for i := 0; i < lenNum; i++ {
			if _, ok := mapString[sSlice[i]]; !ok {
				mapString[sSlice[i]] = i
				total = len(mapString)
			} else {
				end = mapString[sSlice[i]]
				for _, v := range sSlice[start:end+1] {
					delete(mapString, v)
				}
				i--
				if total > max {
					max = total
				}
				total = 0
				start = end + 1
			}
			if i == lenNum - 1 && total > max {
				max = total
			}
		}
		return max
	} else if lenNum==1 {
		return 1
	} else {
		return 0
	}
}

func main()  {
	//s := " "
	//s := "au"
	//s := "abcabcbb"
	//s := "aab"
	//s := "bbbbb"
	//s := "asljlj"
	//s := "dvdf"
	//s := "aabaab!bb"
	s := "bpfbhmipx"
	res := lengthOfLongestSubstring2(s)
	fmt.Println(res)
}