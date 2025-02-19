package main

import "fmt"

func main() {
	fmt.Println(jump([]int{1, 2}))
}

func jump(nums []int) int {
	maxPath := 0
	m := 0
	var count = 0
	for i := 0; i < len(nums); i++ {

		// 如果完成当前这一轮则更新当前的最大值
		if i > maxPath {
			maxPath = m
			count++
		}

		// 保存最大值
		m = max(m, nums[i]+i)
	}

	return count
}
