package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum([]int{1, 4, 6}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var s int //用于记录当期的和
	var traceback func(candidates []int, path []int)

	traceback = func(candidates []int, path []int) {
		switch {
		case s == target:
			p := make([]int, len(path))
			copy(p, path)
			res = append(res, p)
			return
		case s > target:
			return
		case s < target:
			//继续添加
			for i := 0; i < len(candidates); i++ {
				s += candidates[i]
				traceback(candidates[i:], append(path, candidates[i]))
				s -= candidates[i]
			}
		}
	}
	traceback(candidates, []int{})

	return res
}
