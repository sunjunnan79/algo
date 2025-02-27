package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化dp数组，dp[i]表示以nums[i]结尾的子数组的最大和
	dp := make([]int, len(nums))
	dp[0] = nums[0] // 第一个元素的子数组和就是它本身

	maxSum := dp[0] // 初始化最大值为第一个元素

	for i := 1; i < len(nums); i++ {
		// dp[i]是选择继续添加nums[i]到前一个子数组，或者从nums[i]开始新子数组
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i]) // 更新最大值
	}

	return maxSum
}
