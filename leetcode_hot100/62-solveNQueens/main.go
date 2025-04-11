package main

import "fmt"

func main() {
	fmt.Println(len(solveNQueens(4)))
}

func solveNQueens(n int) [][]string {
	var resp [][]string

	var feedback func(temp [][]int, used [][]int, count int)

	feedback = func(temp [][]int, used [][]int, count int) {

		if count == n {
			//转化为最终结果
			var s = make([]string, len(temp))
			for i := 0; i < len(temp); i++ {
				var t string
				for j := 0; j < len(temp[i]); j++ {
					if temp[i][j] == 1 {
						t += "Q"
					} else {
						t += "."
					}
				}
				s[i] = t
			}
			resp = append(resp, s)
			return
		}

		for i := 0; i < n; i++ {
			if !(used[count][i] > 0) {
				//设置皇后
				setUsed(used, count, i, true)
				temp[count][i] = 1
				feedback(temp, used, count+1)
				//回溯
				temp[count][i] = 0
				setUsed(used, count, i, false)
			}

		}
	}

	t := make([][]int, n)
	u := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, n)
		u[i] = make([]int, n)
	}

	feedback(t, u, 0)
	return resp
}

func setUsed(used [][]int, x int, y int, status bool) {
	var n = len(used)
	var count = 0
	if status {
		count = 1
	} else {
		count = -1
	}

	used[x][y] += count

	var setOneRedict = func(a, b int) {

		var x1, y1 = x + a, y + b
		for x1 >= 0 && x1 < n && y1 >= 0 && y1 < n {
			used[x1][y1] += count
			x1 += a
			y1 += b
		}
	}

	setOneRedict(0, 1)
	setOneRedict(0, -1)

	setOneRedict(1, 1)
	setOneRedict(1, 0)
	setOneRedict(1, -1)

	setOneRedict(-1, 0)
	setOneRedict(-1, 1)
	setOneRedict(-1, -1)

}
