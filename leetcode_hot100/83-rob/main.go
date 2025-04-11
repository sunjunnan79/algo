package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	var m int
	for i := 0; i < len(nums); i++ {
		var a, b int
		if i-3 >= 0 {
			a = nums[i-3]
		}
		if i-2 >= 0 {
			b = nums[i-2]
		}
		nums[i] += max(a, b)
		m = max(m, nums[i])
	}
	return m
}
