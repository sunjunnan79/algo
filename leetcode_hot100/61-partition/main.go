package main

import "fmt"

func main() {
	fmt.Println(partition("aacd"))
}

func partition(s string) [][]string {
	var resp [][]string
	var feedback func(used []string, l int)

	feedback = func(used []string, l int) {
		//used表示当前已经用过的字符
		//如果已经全部用完直接添加到结果
		if l == len(s) {
			var str = make([]string, len(used))
			copy(str, used)
			resp = append(resp, str)
			return
		}

		//尝试构造下一个部分,如果是回文则尝试遍历
		for i := l + 1; i <= len(s); i++ {
			//检查是否是回文数,如果是则执行
			if checkHuiwen(s[l:i]) {
				feedback(append(used, s[l:i]), i)
			}
		}

	}
	feedback([]string{}, 0)
	return resp
}
func checkHuiwen(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
