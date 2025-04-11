package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
}

func decodeString(s string) string {
	stackNum := list.New() // 存放倍数
	stackStr := list.New() // 存放待拼接的字符串
	curStr := ""           // 当前正在拼接的字符串
	num := ""              // 当前正在解析的数字

	for _, v := range s {
		switch {
		case v >= '0' && v <= '9': // 数字部分
			num += string(v)
		case v >= 'a' && v <= 'z': // 字母部分
			curStr += string(v)
		case v == '[': // 进入括号
			n, _ := strconv.Atoi(num) // 解析倍数
			stackNum.PushBack(n)      // 存入倍数栈
			stackStr.PushBack(curStr) // 存入字符串栈
			curStr = ""               // 清空当前字符串
			num = ""                  // 清空当前数字
		case v == ']': // 结束括号
			// 获取倍数
			n := stackNum.Back()
			stackNum.Remove(n)

			// 获取前面的字符串
			prevStr := stackStr.Back()
			stackStr.Remove(prevStr)

			// 生成新字符串
			repeatStr := ""
			for i := 0; i < n.Value.(int); i++ {
				repeatStr += curStr
			}
			curStr = prevStr.Value.(string) + repeatStr // 连接前一个字符串
		}
	}
	return curStr
}
