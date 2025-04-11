package main

import "fmt"

func main() {
	fmt.Println(bag(20, [][2]int{
		{2, 3},
		{1, 2},
		{3, 4},
		{2, 2},
		{5, 10},
		{4, 5},
		{1, 3},
		{6, 8},
	}))
}

func bag(maxSize int, items [][2]int) int {
	dp := make([]int, maxSize+1)
	for _, item := range items {
		for i := maxSize; i >= item[0]; i-- {
			dp[i] = max(dp[i], dp[i-item[0]]+item[1])
		}
	}
	return dp[maxSize]
}

//func canPartition(nums []int) bool {
//	sum := 0
//	for _, v := range nums {
//		sum += v
//	}
//	if sum%2 != 0 {
//		return false
//	}
//	target := sum / 2
//	dp := make([]bool, target+1)
//	dp[0] = true
//	for _, num := range nums {
//		for j := target; j >= num; j-- {
//			dp[j] = dp[j] || dp[j-num]
//		}
//	}
//	return dp[target]
//}
