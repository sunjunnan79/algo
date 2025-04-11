package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses("()(()"))
}

func longestValidParentheses(s string) int {
	stack := []int{-1} // 栈初始化，存储起始点
	maxLength := 0     // 记录最大长度

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 遇到 '('，将当前位置的索引压入栈
			stack = append(stack, i)
		} else {
			// 遇到 ')'，先弹出栈顶
			stack = stack[:len(stack)-1]

			// 如果栈非空，计算当前有效括号的长度
			if len(stack) > 0 {
				// 当前有效括号的长度为当前位置 - 栈顶索引
				maxLength = max(maxLength, i-stack[len(stack)-1])
			} else {
				// 如果栈为空，压入当前索引作为新的起始点
				stack = append(stack, i)
			}
		}
	}

	return maxLength
}
