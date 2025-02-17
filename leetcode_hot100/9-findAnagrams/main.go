package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("ababababab", "aab"))
}

func findAnagrams(s string, p string) []int {
	target := make(map[byte]int)
	result := make([]int, 0)

	// 设置目标char的频率
	for i := 0; i < len(p); i++ {
		target[p[i]]++
	}

	// 滑动窗口
	window := make(map[byte]int)
	left, right := 0, 0
	match := 0

	for right < len(s) {
		// 增加右边窗口的字符
		if target[s[right]] > 0 {
			window[s[right]]++
			if window[s[right]] == target[s[right]] {
				match++
			}
		}

		// 当窗口大小达到 p 的长度时，检查窗口是否为异位词
		if right-left+1 == len(p) {
			if match == len(target) {
				result = append(result, left)
			}

			// 移动左边窗口，准备下一次滑动
			if target[s[left]] > 0 {
				if window[s[left]] == target[s[left]] {
					match--
				}
				window[s[left]]--
			}
			left++
		}

		// 扩展右边窗口
		right++
	}

	return result
}
