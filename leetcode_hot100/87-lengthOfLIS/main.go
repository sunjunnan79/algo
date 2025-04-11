package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{3, 10, 4, 3, 8, 9}))
}

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = 1
	m := 1
	for i := 1; i < len(nums); i++ {
		tag := false
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				tag = true
			}
		}

		if !tag {
			dp[i] = 1
		}

		m = max(m, dp[i])
	}
	return m
}
