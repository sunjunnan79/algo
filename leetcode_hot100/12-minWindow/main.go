package main

import "fmt"

func main() {
	fmt.Println(minWindow("a", "a"))
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// 记录 t 中字符的数量
	var letters = make(map[byte]int)
	for i := 0; i < len(t); i++ {
		letters[t[i]]++
	}

	var resp string
	var slow, fast int
	var tempLetters = make(map[byte]int)

	for fast < len(s) {
		// 扩展 fast，直到找出一个包含所有 t 中字符的窗口
		tempLetters[s[fast]]++
		fast++

		// 判断当前窗口是否包含 t 中所有字符
		for compare(tempLetters, letters) {
			// 更新结果
			if fast-slow < len(resp) || len(resp) == 0 {
				resp = s[slow:fast]
			}
			// 缩小窗口
			tempLetters[s[slow]]--
			slow++
		}
	}

	return resp
}

// 比较两个字符计数map
func compare(a, b map[byte]int) bool {
	for k, v := range b {
		if a[k] < v {
			return false
		}
	}
	return true
}
