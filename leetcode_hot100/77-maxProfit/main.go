package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))

}

func maxProfit(prices []int) int {
	maximumProfit, prevMin := 0, prices[0]
	n := len(prices)
	for i := 1; i < n; i++ {
		maximumProfit = max(maximumProfit, prices[i]-prevMin)
		prevMin = min(prevMin, prices[i])
	}
	return maximumProfit
}
