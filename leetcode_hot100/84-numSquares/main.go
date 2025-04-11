package main

import "fmt"

func main() {
	fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	var dp = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		for j := 1; j <= i; j++ {
			if i-j*j < 0 {
				break
			}
			if dp[i] == 0 {
				dp[i] = dp[i-j*j] + 1
			} else {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
	}
	return dp[n]
}
