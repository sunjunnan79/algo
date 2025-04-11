package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	var res []string

	m := n * 2
	path := make([]byte, m)

	var dfs func(int, int)
	dfs = func(i, open int) {
		if i == m {
			res = append(res, string(path))
			return
		}
		if open < n {
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open {
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return res
}
