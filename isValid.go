package main

import "fmt"

func isValid(s string) bool {
	sLen := len(s)
	if sLen ==0 || sLen % 2 == 1 {
		return false
	}
	leftMap := map[string]struct{}{
		"(": {},
		"[":{},
		"{":{},
	}
	rightMap := map[string]string{
		")":"(",
		"]":"[",
		"}":"{",
	}
	var list []string
	for i:=0; i<sLen; i++ {
		temp := s[i:i+1]
		if _,ok := leftMap[temp]; ok {
			list = append(list, temp)
		} else {
			lenList := len(list)
			if lenList == 0 {
				return false
			} else {
				if list[lenList-1] == rightMap[temp] {
					list = list[:lenList-1]
				} else {
					return false
				}
			}
		}
	}
	if len(list) > 0 {
		return false
	}
	return true
}
/**
 范例
 */
func isValid1(s string) bool {

	// 奇数个直接返回false
	if len(s)%2 == 1 {
		return false
	}
	// 辅助栈
	stack := make([]rune, 0)
	// 模式匹配字典 也可使用 switch语句
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	// 遍历字符串的值
	for _, v := range s {
		// val 为栈中需要匹配的元素
		if val, ok := pairs[v]; ok {
			// 栈不为空且栈顶元素不匹配直接结束
			if len(stack) == 0 || val != stack[len(stack)-1] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 不是闭合括号则入栈
			stack = append(stack, v)
		}
	}
	// 如果栈空标识完全匹配
	return len(stack) == 0
}

func main()  {
	//str := "([)]"
	//str := "(("
	str := "()[]"
	res := isValid(str)
	fmt.Print(res)
}