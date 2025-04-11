package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}

func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		if i == 0 {
			dp[i] = grid[0][i]
			continue
		}
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][j]
				continue
			}
			dp[j] = min(dp[j-1]+grid[i][j], dp[j]+grid[i][j])
		}
	}
	return dp[len(dp)-1]
}
