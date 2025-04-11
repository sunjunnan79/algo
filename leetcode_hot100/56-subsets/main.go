package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2}))
}

func subsets(nums []int) [][]int {
	var res [][]int

	var traceback = func(path []int, nums []int) {}

	traceback = func(path []int, nums []int) {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)

		for i := 0; i < len(nums); i++ {
			path = append(path, nums[i])
			traceback(path, nums[i+1:])
			//回溯状态
			path = path[:len(path)-1]

		}
	}

	traceback([]int{}, nums)

	return res
}
